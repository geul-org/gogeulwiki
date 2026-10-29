//ff:func feature=cli type=loader control=sequence
//ff:what TestLanguageRegistry — the newest or a named version of the spec gives its tags and `@n 0x<hash>`; a missing document, an unknown version, a spec without codes and a missing version file fail
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/resolve"
)

func TestLanguageRegistry(t *testing.T) {
	root := t.TempDir()
	if _, _, err := languageRegistry(resolve.Dir{Root: root}, 0); err == nil {
		t.Fatal("no spec document: want error")
	}
	fixtureRegistry(t, root)
	dir := resolve.Dir{Root: root}
	langs, doc, err := languageRegistry(dir, 0)
	if err != nil || len(langs) != 3 || langs["pt-br"] != 0x4130 || langs["zh-hans"] != 0x404C {
		t.Fatalf("%v %v", langs, err)
	}
	h, _ := dir.Log([]core.Word{0}, registryDoc)
	if doc != fmt.Sprintf("0x0000/0x0000000000000848@1 0x%x", h.Entries[0].Hash) {
		t.Fatalf("doc %q", doc)
	}
	if _, doc1, err := languageRegistry(dir, 1); err != nil || doc1 != doc {
		t.Fatalf("version 1: %q %v", doc1, err)
	}
	if _, _, err := languageRegistry(dir, 2); err == nil {
		t.Fatal("version 2: want error")
	}
	empty := t.TempDir()
	fixtureRegistry(t, empty, "none")
	if _, _, err := languageRegistry(resolve.Dir{Root: empty}, 0); err == nil {
		t.Fatal("spec without codes: want error")
	}
	if err := os.Remove(filepath.Join(root, ".well-known/geul/0x0000/0x0000000000000848/0x00000001.geul")); err != nil {
		t.Fatal(err)
	}
	if _, _, err := languageRegistry(dir, 0); err == nil {
		t.Fatal("version file missing: want error")
	}
}
