//ff:func feature=port type=helper control=iteration dimension=1
//ff:what snakSet — snaks as one order-free key: their texts sorted and joined
package port

import (
	"slices"
	"strings"
)

// snakSet is the multiset of texts (snakText) as one string: sorted, so the
// order of the snaks does not count, and duplicates kept.
func snakSet(texts []string) string {
	s := slices.Clone(texts)
	slices.Sort(s)
	var b strings.Builder
	for _, t := range s {
		b.WriteString(t)
		b.WriteByte(';')
	}
	return b.String()
}
