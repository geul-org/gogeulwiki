//ff:func feature=cli type=loader control=sequence
//ff:what TestRegistryDir — a plain tree is read directly, a packed archive through its sets; a broken pack set fails
package main

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
)

func TestRegistryDir(t *testing.T) {
	root := t.TempDir()
	fixtureRegistry(t, root)
	d, err := registryDir(root)
	if err != nil || d.Store != nil {
		t.Fatal("plain tree")
	}
	packed := t.TempDir()
	if _, err := pack.Build(context.Background(), root, filepath.Join(packed, ".well-known/geul/0x0000"), pack.BuildOptions{}); err != nil {
		t.Fatal(err)
	}
	d, err = registryDir(packed)
	if err != nil || d.Store == nil {
		t.Fatalf("packed: %v", err)
	}
	if _, doc, err := languageRegistry(d, 0); err != nil || doc == "" {
		t.Fatalf("registry through the pack: %v", err)
	}
	os.WriteFile(filepath.Join(packed, ".well-known/geul/0x0000/latest.txt"), []byte("1 0000\n"), 0o644)
	if _, err := registryDir(packed); err == nil {
		t.Fatal("broken set: want error")
	}
}
