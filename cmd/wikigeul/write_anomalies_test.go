//ff:func feature=cli type=command control=sequence
//ff:what TestWriteAnomalies — comment head, header, one line per anomaly; an unwritable path fails
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteAnomalies(t *testing.T) {
	p := filepath.Join(t.TempDir(), "a.tsv")
	if err := writeAnomalies(p, "# head", []anomaly{{ID: "Q1"}, {ID: "Q2"}}); err != nil {
		t.Fatal(err)
	}
	b, _ := os.ReadFile(p)
	if string(b) != "# head\n"+anomalyHeader+"\n"+anomalyLine(anomaly{ID: "Q1"})+"\n"+anomalyLine(anomaly{ID: "Q2"})+"\n" {
		t.Fatalf("%q", b)
	}
	if err := writeAnomalies(t.TempDir(), "#", nil); err == nil {
		t.Fatal("folder: want error")
	}
}
