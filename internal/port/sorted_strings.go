//ff:func feature=port type=helper control=sequence
//ff:what sortedStrings — a sorted copy of a string list
package port

import "slices"

// sortedStrings returns a sorted copy of s.
func sortedStrings(s []string) []string {
	out := slices.Clone(s)
	slices.Sort(out)
	return out
}
