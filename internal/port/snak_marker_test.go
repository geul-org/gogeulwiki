//ff:func feature=port type=encoder control=sequence
//ff:what TestSnakMarker — somevalue is UNKNOWN-VALUE, novalue is NO-VALUE, a value snak and an empty type have none
package port

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestSnakMarker(t *testing.T) {
	if r, ok := snakMarker("somevalue"); !ok || !core.IsSlotMarker(r) || r.Codec != uint8(core.UnknownValue) {
		t.Fatalf("somevalue: %+v %v", r, ok)
	}
	if r, ok := snakMarker("novalue"); !ok || !core.IsSlotMarker(r) || r.Codec != uint8(core.NoValue) {
		t.Fatalf("novalue: %+v %v", r, ok)
	}
	if _, ok := snakMarker("value"); ok {
		t.Fatal("value: want false")
	}
	if _, ok := snakMarker(""); ok {
		t.Fatal("empty: want false")
	}
}
