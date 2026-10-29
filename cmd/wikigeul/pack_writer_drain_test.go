//ff:func feature=cli type=command control=sequence
//ff:what TestPackWriterDrain — results in any order are put and the set closed; a failed result and a producer error fail
package main

import (
	"errors"
	"testing"
)

func TestPackWriterDrain(t *testing.T) {
	db, schema := fixtureDB(t)
	c := packTestConfig(t, schema)
	s := packTestStream(t)
	feed := func(rs ...packResult) <-chan packResult {
		ch := make(chan packResult, len(rs))
		for _, r := range rs {
			ch <- r
		}
		close(ch)
		return ch
	}
	a := convertOne(db, c, packTestProv(t, c), s, packJob{Seq: 1, Kind: 'Q', Num: 2})
	b := convertOne(db, c, packTestProv(t, c), s, packJob{Seq: 0, Kind: 'Q', Num: 1})
	var none error
	w, _ := newPackWriter(s, config{Every: 10}, packProgress{})
	if err := w.drain(feed(a, b), &none); err != nil || w.prog.Docs != 2 || w.next != 2 {
		t.Fatal(err)
	}
	bad, _ := newPackWriter(packTestStream(t), config{Every: 10}, packProgress{})
	if err := bad.drain(feed(packResult{Err: errors.New("x")}), &none); err == nil {
		t.Fatal("failed result: want error")
	}
	perr := errors.New("producer")
	p, _ := newPackWriter(packTestStream(t), config{Every: 10}, packProgress{})
	if err := p.drain(feed(), &perr); err != perr {
		t.Fatalf("producer error: %v", err)
	}
}
