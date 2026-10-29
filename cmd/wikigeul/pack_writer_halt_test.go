//ff:func feature=cli type=command control=sequence
//ff:what TestPackWriterHalt — a strict stop checkpoints what was put, writes the report with the stop line and returns an anomalyError; an unwritable report and a failed checkpoint fail differently
package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPackWriterHalt(t *testing.T) {
	path := filepath.Join(t.TempDir(), "a.tsv")
	s := packTestStream(t)
	w, _ := newPackWriter(s, config{Every: 100, Strict: true, Anomaly: path}, packProgress{Kind: 'Q', Num: 3, Docs: 2})
	err := w.halt([]anomaly{{ID: "Q4", Kind: "raw", Reason: "time: 0년"}})
	var ae anomalyError
	if !errors.As(err, &ae) || ae.ID != "Q4" || ae.Count != 1 || ae.Path != path || string(s.Meta()) == "" {
		t.Fatalf("%v", err)
	}
	b, _ := os.ReadFile(path)
	if !strings.HasPrefix(string(b), "# wikigeul -strict: stopped at Q4; checkpoint after Q3, 2 documents\n# wikigeul -from") || !strings.Contains(string(b), "\ttime: 0년") && !strings.Contains(string(b), `"time: 0년"`) {
		t.Fatalf("%q", b)
	}
	w.anomaly = t.TempDir()
	if err := w.halt([]anomaly{{ID: "Q4"}}); err == nil || errors.As(err, &ae) && err.Error() == ae.Error() {
		t.Fatalf("report folder: %v", err)
	}
	w.rep.Close()
	w.w.WriteString("x")
	if err := w.halt([]anomaly{{ID: "Q4"}}); err == nil || !strings.Contains(err.Error(), "checkpoint failed") {
		t.Fatalf("closed report: %v", err)
	}
}
