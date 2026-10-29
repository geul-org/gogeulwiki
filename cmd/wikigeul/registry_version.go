//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what registryVersion — the version list entry to read: version v, or for 0 the newest neither deprecated nor deleted
package main

import "github.com/geul-org/gogeul/pkg/resolve"

// registryVersion picks the entry of h for version v; v 0 is the newest
// version that is neither deprecated nor deleted (as resolve.Dir.At serves
// it). A deleted or unknown version is false.
func registryVersion(h resolve.History, v uint32) (resolve.LogEntry, bool) {
	var pick resolve.LogEntry
	found := false
	for _, e := range h.Entries {
		if v == 0 && !e.Deprecated && !e.Deleted || v != 0 && e.Version == v && !e.Deleted {
			pick, found = e, true
		}
	}
	return pick, found
}
