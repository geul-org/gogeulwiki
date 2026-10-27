//ff:func feature=port type=helper control=selection
//ff:what describedSnak — what one record says of a LOCAL as a reference snak: TRIPLE(P, LOCAL, v) → P=v, REFERENCE-NODE(LOCAL → URI) → P854=URI
package port

import "github.com/geul-org/gogeul/pkg/core"

// describedSnak reads r as a statement about a LOCAL node: a TRIPLE whose
// Subject is a LOCAL gives that LOCAL and snakText(Predicate, Object); a
// REFERENCE-NODE whose Subject is a LOCAL gives snakText(P854, Object). Any
// other record is false.
func describedSnak(r core.Record) (uint64, string, bool) {
	switch r.Codec {
	case edgeTriple:
		_, ps, ok := edgeParts(r, 3)
		s, okS := localOfParts(ps, 1)
		return s, snakText(partOf(ps, 0), partOf(ps, 2)), ok && okS
	case edgeReferenceNode:
		p854, _ := entityRef("P854")
		_, ps, ok := edgeParts(r, 2)
		s, okS := localOfParts(ps, 0)
		return s, snakText(wordsOf(p854), partOf(ps, 1)), ok && okS
	default:
		return 0, "", false
	}
}
