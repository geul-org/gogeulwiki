//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what TestBuildDocMarkers — Q30-shaped claims: a border with P582 somevalue, a novalue central bank with time qualifiers beside a preferred value, novalue and somevalue told apart by a qualifier; hex as the grammar shows (after the dump's REFERENCE-NODE and the HEAD SOURCE IDs), schema clean, counted
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
	"github.com/geul-org/gogeul/pkg/schema"

	"github.com/geul-org/gogeulwiki/internal/source"
)

func TestBuildDocMarkers(t *testing.T) {
	year := func(y string) string {
		return `{"time":"+` + y + `-00-00T00:00:00Z","precision":9,"calendarmodel":"` + gregorian + `"}`
	}
	e := source.Entity{
		ID: "Q30",
		Claims: []source.Claim{
			{Property: "P47", Type: "wikibase-entityid", Value: "Q34266", Rank: "normal", SnakType: "value", Qualifiers: []source.Qualifier{{Property: "P582", Datatype: "time", SnakType: "somevalue"}}},
			{Property: "P1304", Type: "wikibase-entityid", Value: "Q53536", Rank: "preferred", SnakType: "value"},
			{Property: "P1304", Datatype: "wikibase-item", Rank: "normal", SnakType: "novalue", Qualifiers: []source.Qualifier{
				{Property: "P580", Type: "time", Value: year("1811"), SnakType: "value"}, {Property: "P582", Type: "time", Value: year("1816"), SnakType: "value"}}},
			{Property: "P575", Rank: "normal", SnakType: "novalue", Qualifiers: []source.Qualifier{{Property: "P3831", Type: "wikibase-entityid", Value: "Q1", SnakType: "value"}}},
			{Property: "P575", Rank: "normal", SnakType: "somevalue", Qualifiers: []source.Qualifier{{Property: "P3831", Type: "wikibase-entityid", Value: "Q2", SnakType: "value"}}},
		},
	}
	b, rep, err := BuildDoc(e, []core.Word{1, 0}, prov(), freshIDs(1))
	if err != nil || rep.Triples != 5 || rep.Qualifiers != 5 || len(rep.Skipped) != 0 ||
		rep.Markers["UNKNOWN-VALUE"] != 1 || rep.Markers["NO-VALUE"] != 2 || rep.Markers["qualifier UNKNOWN-VALUE"] != 1 {
		t.Fatalf("report %+v %v", rep, err)
	}
	if fs := schema.Check(b, schema.Options{}).Facts; len(fs) != 0 {
		t.Fatalf("schema: %+v", fs)
	}
	d, _ := inspect.Decode(b)
	want := [][]core.Word{
		{0x148A, 0x0088, 5, 0x00A8, 47, 0x00A0, 30, 0x00A0, 34266, 0},
		{0x138B, 0x0088, 6, 0x00A8, 582, 0x0088, 5, core.UnknownValue},
		{0x148A, 0x0088, 7, 0x00A8, 1304, 0x00A0, 30, 0x00A0, 53536, 1},
		{0x140A, 0x0088, 8, 0x00A8, 1304, 0x00A0, 30, core.NoValue, 0},
		{0x148B, 0x0088, 9, 0x00A8, 580, 0x0088, 8, 0x0924, 0x3138, 0x3131},
		{0x148B, 0x0088, 10, 0x00A8, 582, 0x0088, 8, 0x0924, 0x3138, 0x3136},
	}
	if len(d.Body) < len(want)+1 || d.Body[0].Record.Codec != edgeReferenceNode {
		t.Fatalf("body %d", len(d.Body))
	}
	for i, w := range want {
		if !slices.Equal(d.Body[i+1].Words, w) {
			t.Fatalf("record %d: %04X want %04X", i, d.Body[i+1].Words, w)
		}
	}
}
