//ff:func feature=port type=encoder control=sequence
//ff:what TestOriginNode — a description is declared once (TRIPLEs, sorted) and reused in any snak order; a URL-only one is a REFERENCE-NODE
package port

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestOriginNode(t *testing.T) {
	rep := Report{}
	st := newSourcing()
	a := freshIDs(1)
	p854, _ := entityRef("P854")
	p1476, _ := entityRef("P1476")
	uri := core.EncodeBytesValue(codecURIEven, []byte("http://x"))
	title := core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("t"))
	desc := []refSnak{{Property: "P854", Prop: p854, Value: uri}, {Property: "P1476", Prop: p1476, Value: title}}
	o, decl := originNode(a, desc, st, &rep)
	if len(decl) != 2 || decl[0].Codec != edgeTriple || decl[0].Payload[3] != 854 || decl[1].Payload[3] != 1476 || rep.OriginLocals != 1 || rep.OriginEdges != 2 {
		t.Fatalf("node: %+v %+v", o, decl)
	}
	again, decl2 := originNode(a, []refSnak{desc[1], desc[0]}, st, &rep)
	if decl2 != nil || again.Payload[0] != o.Payload[0] || rep.OriginLocals != 1 {
		t.Fatalf("reuse: %+v %v", again, decl2)
	}
	_, decl3 := originNode(a, desc[:1], st, &rep)
	if len(decl3) != 1 || decl3[0].Codec != edgeReferenceNode || rep.OriginLocals != 2 || rep.OriginEdges != 3 {
		t.Fatalf("url: %+v", decl3)
	}
}
