//ff:func feature=port type=encoder control=sequence
//ff:what TestDescribedNode — one TRIPLE per snak on a new LOCAL; rebuilt from the previous version, the node keeps its LOCAL
package port

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/delta"
)

func TestDescribedNode(t *testing.T) {
	p50, _ := entityRef("P50")
	q, _ := entityRef("Q18145749")
	ss := []refSnak{{Property: "P50", Prop: p50, Value: q}}
	rs, local := describedNode(freshIDs(7), ss)
	if len(rs) != 1 || local != 7 || rs[0].Codec != edgeTriple {
		t.Fatalf("%d %+v", local, rs)
	}
	prev := document(nil, rs)
	ids, _ := delta.NewLocalIDs(prev, 20)
	again, l2 := describedNode(ids, ss)
	if l2 != 7 || len(again) != 1 || again[0].Payload[1] != rs[0].Payload[1] {
		t.Fatalf("carried %d %+v", l2, again)
	}
}
