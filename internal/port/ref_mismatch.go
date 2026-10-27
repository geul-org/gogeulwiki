//ff:func feature=port type=helper control=iteration dimension=1
//ff:what refMismatch — the number of Edges whose expected reference snak sets differ from those read back, as multisets
package port

import "slices"

// refMismatch counts the Edge Ids of expect or got whose lists of snak sets
// (one per reference) differ once sorted: reference order is not kept,
// everything else must read back (README 출처, 역변환).
func refMismatch(expect, got map[uint64][]string) int {
	n := 0
	for id, e := range expect {
		if !slices.Equal(sortedStrings(e), sortedStrings(got[id])) {
			n++
		}
	}
	for id := range got {
		if _, ok := expect[id]; !ok {
			n++
		}
	}
	return n
}
