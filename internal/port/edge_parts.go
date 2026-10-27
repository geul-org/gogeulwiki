//ff:func feature=port type=helper control=iteration dimension=1
//ff:what edgeParts — an Edge's LOCAL Edge Id and its next n whole records as words; false when they do not decode
package port

import "github.com/geul-org/gogeul/pkg/core"

// edgeParts splits the payload of the Edge r into its Edge Id (a LOCAL
// Reference) and the n records after it, each as its words. Words after the
// n-th record (TRIPLE's rank) are not returned. It is false when r is not
// an Edge or a part does not decode.
func edgeParts(r core.Record, n int) (uint64, [][]core.Word, bool) {
	lead, k, err := core.DecodeRecord(r.Payload)
	id, lerr := core.DecodeLocalRef(lead)
	if r.Kind != core.KindEdge || len(r.Issuer) > 0 || err != nil || lerr != nil {
		return 0, nil, false
	}
	ws := r.Payload[k:]
	parts := make([][]core.Word, n)
	for i := range parts {
		_, m, err := core.DecodeRecord(ws)
		if err != nil {
			return 0, nil, false
		}
		parts[i], ws = ws[:m], ws[m:]
	}
	return id, parts, true
}
