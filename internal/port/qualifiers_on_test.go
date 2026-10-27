//ff:func feature=port type=helper control=sequence
//ff:what TestQualifiersOn — QUALIFIER snaks grouped by Target in order; other records and broken QUALIFIERs skipped
package port

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestQualifiersOn(t *testing.T) {
	a := freshIDs(1)
	p813, _ := entityRef("P813")
	q, _ := entityRef("Q1")
	broken := core.Record{Kind: core.KindEdge, Codec: edgeQualifier, Payload: []core.Word{0x0088, 9, 0x00A8}}
	got := qualifiersOn([]core.Record{qualifierEdge(a, p813, 5, q), broken, core.EncodeLocalRef(1), qualifierEdge(a, q, 5, p813)})
	if len(got) != 1 || len(got[5]) != 2 || got[5][0] != snakText(wordsOf(p813), wordsOf(q)) {
		t.Fatalf("%v", got)
	}
}
