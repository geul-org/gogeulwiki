//ff:func feature=port type=helper control=iteration dimension=1
//ff:what referenceSnaks — the Wikidata reference snak sets a BODY says of each Edge, read back from its SOURCE Edges (the inverse of sourceEdges)
package port

import (
	"slices"

	"github.com/geul-org/gogeul/pkg/core"
)

// referenceSnaks reads body back into references (README 출처, 역변환): for
// each SOURCE(target, origin), the snak set (snakSet) of the origin's
// description (originSnaks) and of the QUALIFIERs on the SOURCE
// (qualifiersOn), appended to the target's list in document order.
func referenceSnaks(body []core.Record) map[uint64][]string {
	desc, quals := describedBy(body), qualifiersOn(body)
	out := map[uint64][]string{}
	for _, r := range body {
		if r.Codec != edgeSource {
			continue
		}
		id, ps, ok := edgeParts(r, 2)
		t, okT := localOfParts(ps, 0)
		if !ok || !okT {
			continue
		}
		out[t] = append(out[t], snakSet(slices.Concat(originSnaks(ps[1], desc), quals[id])))
	}
	return out
}
