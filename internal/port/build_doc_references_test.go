//ff:func feature=port type=encoder control=sequence
//ff:what TestBuildDocReferences — references become SOURCE Edges that read back to the same snak sets (lossless), pass schema (E8), share equal LOCAL sources, and keep their IDs when rebuilt
package port

import (
	"maps"
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/delta"
	"github.com/geul-org/gogeul/pkg/schema"

	"github.com/geul-org/gogeulwiki/internal/source"
)

func TestBuildDocReferences(t *testing.T) {
	e := source.Entity{ID: "Q42", Claims: []source.Claim{
		{Property: "P19", Type: "wikibase-entityid", Value: "Q350", Rank: "normal", References: q42Refs()},
		{Property: "P20", Type: "wikibase-entityid", Value: "Q159288", Rank: "normal", References: q42Refs()[4:6]},
	}}
	b, rep, err := BuildDoc(e, []core.Word{1, 0}, prov(), freshIDs(1))
	if err != nil || rep.RefMismatch != 0 || rep.Sources != 10 || rep.OriginQIDs != 1 || rep.OriginNone != 2 || rep.OriginLocals != 5 || rep.OriginEdges != 6+7+2+1+1 || rep.SourceQuals != 11 {
		t.Fatalf("report %+v %v", rep, err)
	}
	if fs := schema.Check(b, schema.Options{}).Facts; len(fs) != 0 {
		t.Fatalf("schema: %+v", fs)
	}
	ws, _ := core.ReadWords(b)
	parts, _ := core.SplitDocument(ws)
	got := referenceSnaks(parts.Body)
	keys := slices.Sorted(maps.Keys(got))
	if len(keys) != 2 || len(got[keys[0]]) != 8 || len(got[keys[1]]) != 2 {
		t.Fatalf("read back %v", keys)
	}
	q, _ := entityRef("Q5375741")
	p248, _ := entityRef("P248")
	p1417, _ := entityRef("P1417")
	brit := snakSet([]string{snakText(wordsOf(p248), wordsOf(q)), snakText(wordsOf(p1417), wordsOf(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("biography/Douglas-Adams"))))})
	if !slices.Contains(got[keys[0]], brit) {
		t.Fatalf("Britannica not read back: %v", got[keys[0]])
	}
	ids, _ := delta.NewLocalIDs(b, 0)
	again, rep2, _ := BuildDoc(e, []core.Word{1, 0}, prov(), ids)
	if !slices.Equal(again, b) || rep2.RefMismatch != 0 {
		t.Fatal("rebuild changed the document")
	}
}
