//ff:func feature=cli type=helper control=sequence
//ff:what TestSameRun — progress fields are ignored, any identity field differing is another run
package main

import "testing"

func TestSameRun(t *testing.T) {
	want := packProgress{Commit: "a", RecordedAt: 1}
	if !sameRun(packProgress{Commit: "a", RecordedAt: 1, Kind: 'Q', Num: 5, Docs: 9, Report: 100}, want) {
		t.Fatal("same run")
	}
	if sameRun(packProgress{Commit: "b", RecordedAt: 1}, want) || sameRun(packProgress{Commit: "a", RecordedAt: 2}, want) {
		t.Fatal("other run")
	}
}
