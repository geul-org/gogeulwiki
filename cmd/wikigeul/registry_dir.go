//ff:func feature=cli type=loader control=sequence
//ff:what registryDir — the archive at root to read the language registry from: its pack sets when it has them, else the directory tree
package main

import (
	"context"
	"path/filepath"

	"github.com/geul-org/gogeul/pkg/pack"
	"github.com/geul-org/gogeul/pkg/resolve"
)

// registryDir opens root like the geul CLI's data folder: an archive whose
// .well-known/geul holds pack sets (latest.txt) is read through them
// (pack.OpenDir), a plain tree directly. The published archive is packed.
func registryDir(root string) (resolve.Dir, error) {
	wk := filepath.Join(root, ".well-known", "geul")
	d := resolve.Dir{Root: root}
	if !pack.IsPacked(wk) {
		return d, nil
	}
	arch, err := pack.OpenDir(context.Background(), wk, pack.BuildOptions{})
	d.Store = arch
	return d, err
}
