//ff:func feature=cli type=command control=sequence
//ff:what TestPackWriterFinish — the stream closes into a set; a report that cannot flush and a stream that cannot close fail
package main

import (
	"testing"
	"time"
)

func TestPackWriterFinish(t *testing.T) {
	db, schema := fixtureDB(t)
	c := packTestConfig(t, schema)
	s := packTestStream(t)
	w, _ := newPackWriter(s, config{Every: 10}, packProgress{})
	w.take(convertOne(db, c, packTestProv(t, c), s, packJob{Kind: 'Q', Num: 1}))
	if err := w.finish(time.Now()); err != nil {
		t.Fatal(err)
	}
	empty, _ := newPackWriter(packTestStream(t), config{Every: 10}, packProgress{})
	if err := empty.finish(time.Now()); err == nil {
		t.Fatal("no document: want error")
	}
	bad, _ := newPackWriter(packTestStream(t), config{Every: 10}, packProgress{})
	bad.w.WriteString("x")
	bad.rep.Close()
	if err := bad.finish(time.Now()); err == nil {
		t.Fatal("report that cannot flush: want error")
	}
}
