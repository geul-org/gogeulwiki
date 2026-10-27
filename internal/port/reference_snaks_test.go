//ff:func feature=port type=helper control=sequence
//ff:what TestReferenceSnaks — each SOURCE read back as origin snaks plus its QUALIFIERs, per Target; broken SOURCEs skipped
package port

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestReferenceSnaks(t *testing.T) {
	a := freshIDs(10)
	q, _ := entityRef("Q36578")
	p248, _ := entityRef("P248")
	p227, _ := entityRef("P227")
	v := core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("119033364"))
	src, id := sourceEdge(a, 2, q)
	broken := core.Record{Kind: core.KindEdge, Codec: edgeSource, Payload: []core.Word{0x0088, 30, 0x00A0}}
	got := referenceSnaks([]core.Record{src, qualifierEdge(a, p227, id, v), broken})
	want := snakSet([]string{snakText(wordsOf(p248), wordsOf(q)), snakText(wordsOf(p227), wordsOf(v))})
	if len(got) != 1 || len(got[2]) != 1 || got[2][0] != want {
		t.Fatalf("%v", got)
	}
}
