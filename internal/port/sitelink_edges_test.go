//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what TestSitelinkEdges — about bound once, per page REFERENCE-NODE(LOCAL → URL) and TRIPLE(about, page, item), badges after their page; unknown sites skipped; none gives nothing
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"

	"github.com/geul-org/wikigeul/internal/source"
)

func TestSitelinkEdges(t *testing.T) {
	subject, _ := entityRef("Q42")
	rep := Report{Skipped: map[string]int{}}
	ids := freshIDs(1)
	rs := sitelinkEdges(ids, subject, []source.Sitelink{
		{Site: "xxsite", Title: "x"},
		{Site: "enwiki", Title: "Douglas Adams", Badges: []string{"Q17437796"}},
		{Site: "kowiki", Title: "더글러스 애덤스"},
	}, &rep)
	var codecs []uint8
	for _, r := range rs {
		codecs = append(codecs, r.Codec)
	}
	want := []uint8{edgeReferenceNode, edgeReferenceNode, edgeTriple, edgeReferenceNode, edgeTriple, edgeReferenceNode, edgeTriple}
	if !slices.Equal(codecs, want) || rep.Sitelinks != 2 || rep.Badges != 1 || rep.Skipped["sitelink 사이트: xxsite"] != 1 {
		t.Fatalf("codecs %02X rep %+v", codecs, rep)
	}
	if len(rep.Notes) != 1 || rep.Notes[0] != (Note{Kind: "skipped", Reason: "sitelink 사이트: xxsite", Property: "sitelink", Value: "xxsite x"}) {
		t.Fatalf("notes %+v", rep.Notes)
	}
	// #1 binds LOCAL 2 to schema:about, #3 binds LOCAL 4 to the enwiki page, #5 is TRIPLE(2, 4, Q42).
	if obj, _, err := core.DecodeRecord(rs[0].Payload[4:]); err != nil {
		t.Fatal(err)
	} else if b, _ := core.DecodeBytesValue(obj, codecURIEven); string(b) != "http://schema.org/about" {
		t.Fatalf("about %q", b)
	}
	if obj, _, err := core.DecodeRecord(rs[1].Payload[4:]); err != nil {
		t.Fatal(err)
	} else if b, _ := core.DecodeBytesValue(obj, codecURIEven); string(b) != "https://en.wikipedia.org/wiki/Douglas_Adams" {
		t.Fatalf("page %q", b)
	}
	if !slices.Equal(rs[2].Payload, []core.Word{0x0088, 5, 0x0088, 2, 0x0088, 4, 0x00A0, 42, 0}) {
		t.Fatalf("about triple %04X", rs[2].Payload)
	}
	if rs := sitelinkEdges(ids, subject, nil, &rep); rs != nil {
		t.Fatal("no sitelinks: want none")
	}
}
