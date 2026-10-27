//ff:func feature=port type=helper control=sequence
//ff:what TestLocalOfParts — the LOCAL of an existing part; a missing part is false
package port

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestLocalOfParts(t *testing.T) {
	ps := [][]core.Word{{0x0088, 7}}
	if id, ok := localOfParts(ps, 0); !ok || id != 7 {
		t.Fatalf("%d %v", id, ok)
	}
	if _, ok := localOfParts(ps, 1); ok {
		t.Fatal("missing: want false")
	}
}
