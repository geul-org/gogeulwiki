//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what describedNode — a LOCAL node and one TRIPLE(P, LOCAL, value, normal) per describing snak, the LOCAL carried over by what the TRIPLEs say
package port

import (
	"slices"

	"github.com/geul-org/gogeul/pkg/core"
)

// describedNode asks a for the LOCAL of a node described by ss (IDs.Node:
// the previous version's node with the same Predicate·Object·Rank words keeps
// its number) and returns a TRIPLE(property, LOCAL, value, normal) per snak,
// in the order of ss, with that LOCAL.
func describedNode(a IDs, ss []refSnak) ([]core.Record, uint64) {
	var about [][]core.Word
	for _, s := range ss {
		about = append(about, slices.Concat(wordsOf(s.Prop), wordsOf(s.Value), []core.Word{0}))
	}
	local := a.Node(about)
	var out []core.Record
	for _, s := range ss {
		r, _ := tripleEdge(a, s.Prop, core.EncodeLocalRef(local), s.Value, 0)
		out = append(out, r)
	}
	return out, local
}
