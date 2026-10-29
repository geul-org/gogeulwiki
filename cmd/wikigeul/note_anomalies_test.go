//ff:func feature=cli type=command control=sequence
//ff:what TestNoteAnomalies — normal and allowed notes pass; the rest become anomalies with property, value, reason and class
package main

import (
	"testing"

	"github.com/geul-org/wikigeul/internal/port"
)

func TestNoteAnomalies(t *testing.T) {
	notes := []port.Note{
		{Kind: "skipped", Reason: "qualifier 개체: L", Property: "P1", Value: "L1"},
		{Kind: "raw", Reason: "reference time: 0년", Property: "P2", Value: "{}"},
		{Kind: "raw", Reason: "quantity: 단위", Property: "P3", Value: "{q}"},
	}
	got := noteAnomalies("Q9", notes, allowSet([]string{"raw:time: 0년"}))
	if len(got) != 1 || got[0] != (anomaly{ID: "Q9", Kind: "raw", Property: "P3", Value: "{q}", Reason: "quantity: 단위", Class: "문법 부족"}) {
		t.Fatalf("%+v", got)
	}
}
