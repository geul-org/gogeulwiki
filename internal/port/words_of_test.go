//ff:func feature=port type=helper control=sequence
//ff:what TestWordsOf — a LOCAL Reference as its two words; a record that does not encode gives none
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestWordsOf(t *testing.T) {
	if got := wordsOf(core.EncodeLocalRef(5)); !slices.Equal(got, []core.Word{0x0088, 5}) {
		t.Fatalf("%04x", got)
	}
	if got := wordsOf(core.Record{Kind: core.KindValue, Codec: 0x80}); len(got) != 0 {
		t.Fatalf("bad codec: %d", len(got))
	}
}
