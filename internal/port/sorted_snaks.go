//ff:func feature=port type=helper control=sequence
//ff:what sortedSnaks — snaks ordered by property number, values of one property kept in their order
package port

import (
	"cmp"
	"slices"
)

// sortedSnaks returns ss sorted by property number (P50 before P248), stable
// so the values of one property (several authors) keep their order. Two
// references with the same description in another snaks-order then describe
// the same LOCAL source with the same bytes.
func sortedSnaks(ss []refSnak) []refSnak {
	out := slices.Clone(ss)
	slices.SortStableFunc(out, func(a, b refSnak) int {
		_, x, _ := entityNumber(a.Property)
		_, y, _ := entityNumber(b.Property)
		return cmp.Compare(x, y)
	})
	return out
}
