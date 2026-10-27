//ff:func feature=port type=encoder control=sequence
//ff:what TestLocalDefinition — REFERENCE-NODE(edge, LOCAL → object) with two fresh IDs; the same object in a later version keeps them
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/delta"
)

func TestLocalDefinition(t *testing.T) {
	obj := core.EncodeBytesValue(codecURIEven, []byte("http://schema.org/about"))
	ids := freshIDs(3)
	r, local := localDefinition(ids, obj)
	objWords, _ := core.EncodeRecord(obj)
	if r.Kind != core.KindEdge || r.Codec != edgeReferenceNode || local != 4 || !slices.Equal(r.Payload, append([]core.Word{0x0088, 3, 0x0088, 4}, objWords...)) {
		t.Fatalf("got %+v local %d", r, local)
	}
	prev := document(nil, []core.Record{r})
	later, err := delta.NewLocalIDs(prev, 9)
	if err != nil {
		t.Fatal(err)
	}
	if again, l := localDefinition(later, obj); l != 4 || !slices.Equal(again.Payload, r.Payload) {
		t.Fatalf("carried over: %+v %d", again, l)
	}
}
