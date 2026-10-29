//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what allowSet — the -allow entries ("kind:reason") as a set
package main

// allowSet makes a set of the -allow entries.
func allowSet(entries []string) map[string]bool {
	out := map[string]bool{}
	for _, e := range entries {
		out[e] = true
	}
	return out
}
