//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what TestRegistryVersion — 0 picks the newest neither deprecated nor deleted; a named version is found unless deleted or unknown
package main

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/resolve"
)

func TestRegistryVersion(t *testing.T) {
	entry := func(v uint32, dep, del bool) resolve.LogEntry {
		return resolve.LogEntry{VersionEntry: resolve.VersionEntry{Version: v, Deprecated: dep, Deleted: del}}
	}
	h := resolve.History{Entries: []resolve.LogEntry{entry(1, false, false), entry(2, false, false), entry(3, true, false), entry(4, false, true)}}
	for v, want := range map[uint32]uint32{0: 2, 1: 1, 3: 3} {
		if e, ok := registryVersion(h, v); !ok || e.Version != want {
			t.Errorf("v%d: %+v %v", v, e, ok)
		}
	}
	for _, v := range []uint32{4, 5} {
		if _, ok := registryVersion(h, v); ok {
			t.Errorf("v%d: want false", v)
		}
	}
	if _, ok := registryVersion(resolve.History{}, 0); ok {
		t.Error("empty: want false")
	}
}
