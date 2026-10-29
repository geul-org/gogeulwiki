//ff:func feature=cli type=command control=sequence
//ff:what TestOpenPackStream — a new stream; the same run resumes at its checkpoint (RECORDED-AT from it when not given); another run, no stream, a checkpoint that is not progress and failing samples are refused
package main

import (
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"
)

func TestOpenPackStream(t *testing.T) {
	db, schema := fixtureDB(t)
	c := packTestConfig(t, schema)
	set := filepath.Join(c.Root, "set")
	s, p, err := openPackStream(set, c, packTestProv(t, c), db, nil, "")
	if err != nil || p.RecordedAt != c.RecordedAt || p.Commit != c.Commit {
		t.Fatalf("%v", err)
	}
	p.Kind, p.Num, p.Docs = 'Q', 3, 2
	meta, _ := json.Marshal(p)
	s.Checkpoint(meta)
	c.Resume = true
	if _, got, err := openPackStream(set, c, packTestProv(t, c), db, nil, ""); err != nil || got.Num != 3 || got.Docs != 2 {
		t.Fatalf("resume: %v", err)
	}
	c0 := c
	c0.RecordedAt = 0
	if _, got, err := openPackStream(set, c0, packTestProv(t, c0), db, nil, ""); err != nil || got.RecordedAt != c.RecordedAt {
		t.Fatalf("resume without -recorded-at: %v", err)
	}
	other := c
	other.Commit = strings.Repeat("b", 40)
	if _, _, err := openPackStream(set, other, packTestProv(t, other), db, nil, ""); err == nil {
		t.Fatal("other run: want error")
	}
	if _, _, err := openPackStream(filepath.Join(c.Root, "none"), c, packTestProv(t, c), db, nil, ""); err == nil {
		t.Fatal("no stream: want error")
	}
	s.Checkpoint([]byte("not json"))
	if _, _, err := openPackStream(set, c, packTestProv(t, c), db, nil, ""); err == nil {
		t.Fatal("not progress: want error")
	}
	bad := packTestConfig(t, "1x")
	if _, _, err := openPackStream(filepath.Join(bad.Root, "set"), bad, packTestProv(t, bad), db, nil, ""); err == nil {
		t.Fatal("failing samples: want error")
	}
}
