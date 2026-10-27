//ff:func feature=port type=helper control=iteration dimension=1
//ff:what qualifiersOn — for each Edge Id of a BODY, the snaks of the QUALIFIERs on it (Property=Value), in order
package port

import "github.com/geul-org/gogeul/pkg/core"

// qualifiersOn maps each Target Edge Id of the QUALIFIERs of body to their
// snakText(Property, Value), in document order.
func qualifiersOn(body []core.Record) map[uint64][]string {
	out := map[uint64][]string{}
	for _, r := range body {
		if r.Codec != edgeQualifier {
			continue
		}
		_, ps, ok := edgeParts(r, 3)
		if t, okT := localOfParts(ps, 1); ok && okT {
			out[t] = append(out[t], snakText(ps[0], ps[2]))
		}
	}
	return out
}
