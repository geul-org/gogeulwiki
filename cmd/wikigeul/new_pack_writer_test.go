//ff:func feature=cli type=command control=sequence
//ff:what TestNewPackWriter — a new run writes the run line and header; a resumed one cuts the report to the checkpoint; no report path writes nowhere; an unopenable report and one that cannot be cut fail
package main

import (
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
)

func TestNewPackWriter(t *testing.T) {
	c := config{Report: filepath.Join(t.TempDir(), "r.tsv"), Every: 3}
	w, err := newPackWriter(packTestStream(t), c, packProgress{})
	if err != nil || w.every != 3 {
		t.Fatal(err)
	}
	w.w.Flush()
	b, _ := os.ReadFile(c.Report)
	if !strings.HasPrefix(string(b), "# wikigeul") || !strings.Contains(string(b), reportHeader) || w.prog.Report != int64(len(b)) {
		t.Fatalf("%q %d", b, w.prog.Report)
	}
	os.WriteFile(c.Report, []byte("0123456789"), 0o644)
	if _, err := newPackWriter(packTestStream(t), c, packProgress{Docs: 1, Kind: 'Q', Report: 4}); err != nil {
		t.Fatal(err)
	}
	if b, _ := os.ReadFile(c.Report); string(b) != "0123" {
		t.Fatalf("cut: %q", b)
	}
	if _, err := newPackWriter(packTestStream(t), config{}, packProgress{Docs: 1, Kind: 'Q', Report: 4}); err != nil {
		t.Fatal("no report on resume")
	}
	if _, err := newPackWriter(packTestStream(t), config{Report: t.TempDir()}, packProgress{}); err == nil {
		t.Fatal("folder as report: want error")
	}
	fifo := filepath.Join(t.TempDir(), "fifo")
	if err := syscall.Mkfifo(fifo, 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := newPackWriter(packTestStream(t), config{Report: fifo}, packProgress{}); err == nil {
		t.Fatal("report that cannot be cut (FIFO): want error")
	}
	bad := filepath.Join(t.TempDir(), "r")
	os.WriteFile(bad, nil, 0o444)
	if _, err := newPackWriter(packTestStream(t), config{Report: bad}, packProgress{}); err == nil {
		t.Fatal("read-only report: want error")
	}
}
