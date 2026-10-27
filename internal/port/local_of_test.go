//ff:func feature=port type=helper control=sequence
//ff:what TestLocalOf — a LOCAL Reference's ID; another Reference or broken words are false
package port

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestLocalOf(t *testing.T) {
	if id, ok := localOf([]core.Word{0x0088, 7}); !ok || id != 7 {
		t.Fatalf("%d %v", id, ok)
	}
	if _, ok := localOf([]core.Word{0x00A0, 7}); ok {
		t.Fatal("QID: want false")
	}
	if _, ok := localOf([]core.Word{0x0089}); ok {
		t.Fatal("broken: want false")
	}
}
