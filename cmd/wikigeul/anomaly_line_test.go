//ff:func feature=cli type=helper control=sequence
//ff:what TestAnomalyLine — six tab-separated columns, value and reason quoted (tabs, newlines), a long value cut with its size
package main

import (
	"strings"
	"testing"
)

func TestAnomalyLine(t *testing.T) {
	got := anomalyLine(anomaly{ID: "Q1", Kind: "raw", Property: "P585", Value: "a\tb", Reason: "x\ny", Class: "기타"})
	if got != "Q1\traw\tP585\t\"a\\tb\"\t\"x\\ny\"\t기타" {
		t.Fatalf("%q", got)
	}
	long := anomalyLine(anomaly{Value: strings.Repeat("v", 5000)})
	if !strings.Contains(long, "…(5000 bytes)") || strings.Count(long, "v") != 4096 {
		t.Fatalf("long: %d", len(long))
	}
}
