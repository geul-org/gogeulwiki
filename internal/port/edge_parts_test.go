//ff:func feature=port type=helper control=iteration dimension=1
//ff:what TestEdgeParts — Edge Id and n records split (rank left out); not an Edge, an ISSUER Edge, a bad Edge Id or a short payload are false
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestEdgeParts(t *testing.T) {
	id, ps, ok := edgeParts(core.Record{Kind: core.KindEdge, Codec: edgeTriple, Payload: []core.Word{0x0088, 4, 0x00A8, 31, 0x0088, 20, 0x00A0, 5, 0}}, 3)
	if !ok || id != 4 || len(ps) != 3 || !slices.Equal(ps[1], []core.Word{0x0088, 20}) || !slices.Equal(ps[2], []core.Word{0x00A0, 5}) {
		t.Fatalf("%d %v %v", id, ps, ok)
	}
	for i, r := range []core.Record{
		core.EncodeLocalRef(3),
		{Kind: core.KindEdge, Codec: edgeTriple, Issuer: []core.Word{1}, Payload: []core.Word{0x0088, 4}},
		{Kind: core.KindEdge, Codec: edgeTriple, Payload: []core.Word{0x00A0, 4}},
		{Kind: core.KindEdge, Codec: edgeTriple, Payload: []core.Word{0x0088, 4, 0x00A8, 31}},
	} {
		if _, _, ok := edgeParts(r, 3); ok {
			t.Fatalf("case %d: want false", i)
		}
	}
}
