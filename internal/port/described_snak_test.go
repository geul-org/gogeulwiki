//ff:func feature=port type=helper control=sequence
//ff:what TestDescribedSnak — TRIPLE on a LOCAL gives P=value, REFERENCE-NODE gives P854=URI; a TRIPLE on an item and other codecs are false
package port

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDescribedSnak(t *testing.T) {
	a := freshIDs(1)
	p50, _ := entityRef("P50")
	p854, _ := entityRef("P854")
	q, _ := entityRef("Q5")
	uri := core.EncodeBytesValue(codecURIEven, []byte("http://x"))
	tr, _ := tripleEdge(a, p50, core.EncodeLocalRef(20), q, 0)
	if l, s, ok := describedSnak(tr); !ok || l != 20 || s != snakText(wordsOf(p50), wordsOf(q)) {
		t.Fatalf("triple %d %q %v", l, s, ok)
	}
	def, local := localDefinition(a, uri)
	if l, s, ok := describedSnak(def); !ok || l != local || s != snakText(wordsOf(p854), wordsOf(uri)) {
		t.Fatalf("refnode %d %q %v", l, s, ok)
	}
	onItem, _ := tripleEdge(a, p50, q, q, 0)
	if _, _, ok := describedSnak(onItem); ok {
		t.Fatal("item subject: want false")
	}
	if _, _, ok := describedSnak(qualifierEdge(a, p50, 1, q)); ok {
		t.Fatal("qualifier: want false")
	}
}
