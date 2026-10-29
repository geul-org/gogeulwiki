//ff:func feature=cli type=command control=sequence
//ff:what TestPackWriterPut — a document goes in as list and version, counted, checkpointing when due; a missing entity only reports; an error, another TRANSFORMED-BY and a key out of order fail
package main

import (
	"errors"
	"strings"
	"testing"
)

func TestPackWriterPut(t *testing.T) {
	db, schema := fixtureDB(t)
	c := packTestConfig(t, schema)
	s := packTestStream(t)
	w, _ := newPackWriter(s, config{Every: 1}, packProgress{})
	conv := func(c config, n uint64) packResult {
		return convertOne(db, c, packTestProv(t, c), s, packJob{Kind: 'Q', Num: n})
	}
	if err := w.put(conv(c, 1)); err != nil || w.prog.Docs != 1 || s.Count() != 2 || string(s.Meta()) == "" {
		t.Fatalf("%v", err)
	}
	if err := w.put(packResult{Job: packJob{Kind: 'Q', Num: 5}, Missing: true}); err != nil || w.prog.Num != 5 || w.prog.Docs != 1 {
		t.Fatal("missing")
	}
	other := c
	other.Commit = strings.Repeat("b", 40)
	if err := w.put(conv(other, 2)); err == nil {
		t.Fatal("other TRANSFORMED-BY: want error")
	}
	if err := w.put(conv(c, 1)); err == nil {
		t.Fatal("same document again: want error")
	}
	if err := w.put(packResult{Job: packJob{Kind: 'Q', Num: 9}, Err: errors.New("x")}); err == nil {
		t.Fatal("failed conversion: want error")
	}
	lazy, _ := newPackWriter(s, config{Every: 10}, packProgress{})
	if err := lazy.put(conv(c, 7)); err != nil || lazy.since != 1 {
		t.Fatalf("no checkpoint yet: %v", err)
	}
}
