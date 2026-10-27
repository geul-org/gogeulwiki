//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what citeEdges — a QUALIFIER on a SOURCE Edge for each citing snak, in order
package port

import "github.com/geul-org/gogeul/pkg/core"

// citeEdges writes each snak of cite (retrieved, imported from, page,
// identifiers …) as QUALIFIER(property, source, value) on the SOURCE Edge
// source, in the order of cite.
func citeEdges(a IDs, source uint64, cite []refSnak) []core.Record {
	var out []core.Record
	for _, s := range cite {
		out = append(out, qualifierEdge(a, s.Prop, source, s.Value))
	}
	return out
}
