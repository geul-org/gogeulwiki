//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what sourceEdges — per reference of a TRIPLE: the Origin's declaration the first time, a SOURCE on the TRIPLE and QUALIFIERs for the citing snaks; the snak set noted for the round trip
package port

import (
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/wikigeul/internal/source"
)

// sourceEdges writes the references refs of the TRIPLE target in order
// (README 출처): each reference's snaks are converted (refSnaks, languages by
// langs) and split
// into those describing the source and those describing the citation
// (splitSnaks); the Origin comes from the first (originOf, declared before
// the SOURCE when new), the citing snaks become QUALIFIERs on the SOURCE
// (citeEdges). The converted snak set of each reference is noted in
// st.expect[target] so BuildDoc can check it reads back (referenceSnaks).
func sourceEdges(a IDs, target uint64, refs []source.Reference, st *sourcing, langs Languages, rep *Report) []core.Record {
	var out []core.Record
	for _, r := range refs {
		ss := refSnaks(r, langs, rep)
		describe, cite := splitSnaks(ss)
		origin, decl := originOf(a, describe, st, rep)
		src, id := sourceEdge(a, target, origin)
		out = append(append(append(out, decl...), src), citeEdges(a, id, cite)...)
		st.expect[target] = append(st.expect[target], snakSet(snakTexts(ss)))
		rep.Sources++
		rep.SourceQuals += len(cite)
	}
	return out
}
