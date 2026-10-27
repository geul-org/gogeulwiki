//ff:func feature=port type=helper control=sequence
//ff:what TestDescribedBy — snaks grouped by LOCAL in order; other records skipped
package port

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDescribedBy(t *testing.T) {
	a := freshIDs(1)
	p50, _ := entityRef("P50")
	q1, _ := entityRef("Q1")
	q2, _ := entityRef("Q2")
	t1, _ := tripleEdge(a, p50, core.EncodeLocalRef(20), q1, 0)
	t2, _ := tripleEdge(a, p50, core.EncodeLocalRef(20), q2, 0)
	got := describedBy([]core.Record{t1, core.EncodeLocalRef(3), t2})
	if len(got) != 1 || len(got[20]) != 2 || got[20][1] != snakText(wordsOf(p50), wordsOf(q2)) {
		t.Fatalf("%v", got)
	}
}
