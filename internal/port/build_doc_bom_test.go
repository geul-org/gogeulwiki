//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what TestBuildDocBOM — texts starting with U+FEFF (label, monolingual and string values, qualifier, reference) are written as they are, pass schema and read back
package port

import (
	"bytes"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/schema"

	"github.com/geul-org/wikigeul/internal/source"
)

func TestBuildDocBOM(t *testing.T) {
	e := source.Entity{
		ID:     "Q100262271",
		Labels: []source.Text{{Lang: "en", Value: "\ufeffLabel"}},
		Claims: []source.Claim{
			{Property: "P1476", Type: "monolingualtext", Value: `{"text":"\ufeffTitle","language":"pt"}`, Rank: "normal",
				Qualifiers: []source.Qualifier{{Property: "P1932", Type: "string", Datatype: "string", Value: "\ufeffAbecásia (GE)"}},
				References: []source.Reference{{Snaks: []source.Qualifier{{Property: "P1476", Type: "monolingualtext", Value: `{"text":"\ufeffSource","language":"en"}`}}}}},
			{Property: "P2093", Type: "string", Datatype: "string", Value: "\ufeffAuthor", Rank: "normal"},
		},
	}
	b, rep, err := BuildDoc(e, []core.Word{1, 0}, prov(), freshIDs(1), nil)
	if err != nil || rep.Labels != 1 || rep.Triples != 2 || rep.Qualifiers != 1 || rep.RefMismatch != 0 || len(rep.Skipped) != 0 || len(rep.Raw) != 0 {
		t.Fatalf("report %+v %v", rep, err)
	}
	if fs := schema.Check(b, schema.Options{}).Facts; len(fs) != 0 {
		t.Fatalf("schema: %+v", fs)
	}
	for _, s := range []string{"\ufeffLabel", "\ufeffTitle", "\ufeffAbecásia (GE)", "\ufeffSource", "\ufeffAuthor"} {
		if !bytes.Contains(b, []byte(s)) {
			t.Errorf("%q not kept", s)
		}
	}
}
