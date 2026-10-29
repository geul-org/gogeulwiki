//ff:func feature=port type=encoder control=sequence
//ff:what TestTripleEdges — convertible claims become TRIPLEs, each followed by its QUALIFIERs; bad predicates, values and ranks are counted, raw-kept dates too; somevalue/novalue become slot markers and are counted; properties used are collected
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/wikigeul/internal/source"
)

func TestTripleEdges(t *testing.T) {
	subj, _ := entityRef("Q1")
	rep := Report{Skipped: map[string]int{}, Raw: map[string]int{}}
	ids := freshIDs(1)
	rs := tripleEdges(ids, subj, []source.Claim{
		{Property: "P31", Type: "wikibase-entityid", Value: "Q5", Rank: "normal", Qualifiers: []source.Qualifier{{Property: "P580", Type: "time", Value: "x"}}},
		{Property: "P1659", Type: "wikibase-entityid", Value: "P279", Rank: "preferred"},
		{Property: "X1", Type: "string", Value: "a", Rank: "normal"},
		{Property: "P2", Type: "unknown", Value: "{}", Rank: "normal"},
		{Property: "P3", Type: "string", Value: "a", Rank: "odd"},
		{Property: "P585", Type: "time", Value: "x", Rank: "normal"},
		{Property: "P40", SnakType: "novalue", Rank: "normal", Qualifiers: []source.Qualifier{{Property: "P582", SnakType: "somevalue"}}},
	}, newSourcing(), nil, &rep)
	if rep.Raw["time: JSON 아님"] != 1 || rep.Raw["qualifier time: JSON 아님"] != 1 || rep.Triples != 4 || rep.Qualifiers != 2 || rep.Markers["NO-VALUE"] != 1 || rep.Markers["qualifier UNKNOWN-VALUE"] != 1 {
		t.Fatalf("raw %v", rep.Raw)
	}
	if len(rs) != 6 || ids.Next() != 7 || rs[4].Payload[len(rs[4].Payload)-2] != 0x0002 || rs[5].Payload[len(rs[5].Payload)-1] != 0x0001 || rs[0].Codec != edgeTriple || rs[1].Codec != edgeQualifier || rs[1].Payload[5] != 1 || rs[2].Payload[len(rs[2].Payload)-1] != 1 {
		t.Fatalf("got %d next %d", len(rs), ids.Next())
	}
	if rep.Skipped["속성 ID: X1"] != 1 || rep.Skipped["형식 미지원: unknown"] != 1 || rep.Skipped["rank: odd"] != 1 {
		t.Fatalf("skipped %v", rep.Skipped)
	}
	if len(rep.Notes) != 5 || rep.Notes[0] != (Note{Kind: "raw", Reason: "qualifier time: JSON 아님", Property: "P580", Value: "x"}) || rep.Notes[1] != (Note{Kind: "skipped", Reason: "속성 ID: X1", Property: "X1", Value: "a"}) || rep.Notes[4] != (Note{Kind: "raw", Reason: "time: JSON 아님", Property: "P585", Value: "x"}) {
		t.Fatalf("notes %+v", rep.Notes)
	}
	if !slices.Equal(rep.Properties, []string{"P31", "P580", "P1659", "P279", "P585", "P40", "P582"}) {
		t.Fatalf("properties %v", rep.Properties)
	}
}
