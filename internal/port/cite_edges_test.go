//ff:func feature=port type=encoder control=sequence
//ff:what TestCiteEdges — one QUALIFIER per citing snak on the SOURCE, in order; none for none
package port

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCiteEdges(t *testing.T) {
	p227, _ := entityRef("P227")
	p813, _ := entityRef("P813")
	v := core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("1"))
	rs := citeEdges(freshIDs(1), 9, []refSnak{{Prop: p227, Value: v}, {Prop: p813, Value: v}})
	if len(rs) != 2 || rs[0].Codec != edgeQualifier || rs[0].Payload[3] != 227 || rs[0].Payload[5] != 9 || rs[1].Payload[3] != 813 {
		t.Fatalf("%+v", rs)
	}
	if citeEdges(freshIDs(1), 9, nil) != nil {
		t.Fatal("none: want nil")
	}
}
