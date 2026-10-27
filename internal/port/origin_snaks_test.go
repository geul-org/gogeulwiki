//ff:func feature=port type=helper control=sequence
//ff:what TestOriginSnaks — an item is P248=item, a LOCAL its description, ABSENT and UNKNOWN-VALUE nothing
package port

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestOriginSnaks(t *testing.T) {
	q, _ := entityRef("Q36578")
	p248, _ := entityRef("P248")
	desc := map[uint64][]string{20: {"a", "b"}}
	if got := originSnaks(wordsOf(q), desc); len(got) != 1 || got[0] != snakText(wordsOf(p248), wordsOf(q)) {
		t.Fatalf("item %v", got)
	}
	if got := originSnaks([]core.Word{0x0088, 20}, desc); len(got) != 2 {
		t.Fatalf("local %v", got)
	}
	if got := originSnaks([]core.Word{core.Absent}, desc); got != nil {
		t.Fatalf("absent %v", got)
	}
	if got := originSnaks([]core.Word{core.UnknownValue}, desc); got != nil {
		t.Fatalf("unknown %v", got)
	}
}
