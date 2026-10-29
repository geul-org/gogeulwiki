//ff:func feature=cli type=command control=sequence
//ff:what TestPackWriterLogProgress — the progress line has the time, last entity, run and session documents, rate and reason counts
package main

import (
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

func TestPackWriterLogProgress(t *testing.T) {
	start := time.Unix(1_800_000_000, 0)
	w := &packWriter{prog: packProgress{Kind: 'Q', Num: 42, Docs: 30}, docs0: 10, start: start, skipped: map[string]int{"개체: L": 2}, raw: map[string]int{}}
	old := os.Stdout
	r, wr, _ := os.Pipe()
	os.Stdout = wr
	w.logProgress(start.Add(10 * time.Second))
	w.logProgress(start)
	wr.Close()
	os.Stdout = old
	b, _ := io.ReadAll(r)
	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	if len(lines) != 2 || !strings.HasPrefix(lines[0], "progress ") || !strings.Contains(lines[0], " last=Q42 docs=30 session_docs=20 rate=2.0 skipped={개체: L=2} raw={}") || !strings.Contains(lines[1], "session_docs=20 rate=") {
		t.Fatalf("%q", b)
	}
}
