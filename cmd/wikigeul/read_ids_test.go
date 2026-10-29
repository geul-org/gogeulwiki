//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what TestReadIDs — IDs sorted P before Q and by number, duplicates and comments dropped, with the file hash; no path is no list; a missing file and a bad ID fail
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadIDs(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "ids")
	os.WriteFile(p, []byte("Q42\n# c\n\nP31\nQ5\nQ42\n"), 0o644)
	js, h, err := readIDs(p)
	if err != nil || len(h) != 64 || len(js) != 3 || js[0] != (packJob{Kind: 'P', Num: 31}) || js[1].Num != 5 || js[2].Num != 42 {
		t.Fatalf("%v %v", js, err)
	}
	if js, h, err := readIDs(""); js != nil || h != "" || err != nil {
		t.Fatal("no list")
	}
	if _, _, err := readIDs(filepath.Join(dir, "none")); err == nil {
		t.Fatal("missing: want error")
	}
	for _, bad := range []string{"L1\n", "Qx\n", "Q0\n"} {
		os.WriteFile(p, []byte(bad), 0o644)
		if _, _, err := readIDs(p); err == nil {
			t.Errorf("%q: want error", bad)
		}
	}
}
