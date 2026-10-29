//ff:func feature=port type=encoder control=sequence
//ff:what TestBuildDocLanguages — with a registry, numbered labels and monolingual values get LANGUAGE-SHORT-CODE, unnumbered ones (private use, unregistered) LANGUAGE-BCP47-CODE; the document passes schema and its references still read back
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/schema"

	"github.com/geul-org/wikigeul/internal/source"
)

func TestBuildDocLanguages(t *testing.T) {
	e := source.Entity{
		ID:     "Q42",
		Labels: []source.Text{{Lang: "en", Value: "Douglas Adams"}, {Lang: "pt-br", Value: "Douglas Adams"}, {Lang: "de-formal", Value: "Douglas Adams"}, {Lang: "en-uk", Value: "Douglas Adams"}},
		Claims: []source.Claim{{Property: "P1559", Type: "monolingualtext", Value: `{"text":"Douglas Adams","language":"en"}`, Rank: "normal",
			References: []source.Reference{{Snaks: []source.Qualifier{{Property: "P1476", Type: "monolingualtext", Value: `{"text":"Obituary","language":"en"}`}}}}}},
	}
	langs := Languages{"en": 0x0784, "pt-br": 0x4130}
	b, rep, err := BuildDoc(e, []core.Word{1, 0}, prov(), freshIDs(1), langs)
	if err != nil || rep.Labels != 4 || rep.Triples != 1 || rep.RefMismatch != 0 {
		t.Fatalf("report %+v %v", rep, err)
	}
	if fs := schema.Check(b, schema.Options{}).Facts; len(fs) != 0 {
		t.Fatalf("schema: %+v", fs)
	}
	ws, _ := core.ReadWords(b)
	count := func(seq ...core.Word) int {
		n := 0
		for i := range ws {
			if i+len(seq) <= len(ws) && slices.Equal(ws[i:i+len(seq)], seq) {
				n++
			}
		}
		return n
	}
	if n := count(0x08C8, 0x0784); n != 3 {
		t.Errorf("en as SHORT-CODE: %d, want 3 (label, value, reference)", n)
	}
	if n := count(0x08C8, 0x4130); n != 1 {
		t.Errorf("pt-BR as SHORT-CODE: %d", n)
	}
	if n := count(0x09C7, 0x656E, 0x2D55, 0x4B00); n != 1 { // "en-UK", canonical case, LANGUAGE-BCP47-CODE-ODD
		t.Errorf("en-uk as BCP 47 en-UK: %d", n)
	}
}
