//ff:func feature=cli type=command control=sequence
//ff:what TestPackWriterTake — results wait for their turn and go out in job order; a failing put stops it
package main

import (
	"errors"
	"testing"
)

func TestPackWriterTake(t *testing.T) {
	missing := func(seq int64) packResult {
		return packResult{Job: packJob{Seq: seq, Kind: 'Q', Num: uint64(seq)}, Missing: true}
	}
	w, _ := newPackWriter(packTestStream(t), config{Every: 10}, packProgress{})
	if err := w.take(missing(1)); err != nil || w.next != 0 || len(w.pending) != 1 {
		t.Fatal("waits for 0")
	}
	if err := w.take(missing(0)); err != nil || w.next != 2 || len(w.pending) != 0 || w.prog.Num != 1 {
		t.Fatalf("drained: %d %v", w.next, err)
	}
	if err := w.take(packResult{Job: packJob{Seq: 2}, Err: errors.New("x")}); err == nil {
		t.Fatal("failing result: want error")
	}
}
