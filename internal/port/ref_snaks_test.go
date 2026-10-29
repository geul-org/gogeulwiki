//ff:func feature=port type=encoder control=sequence
//ff:what TestRefSnaks — convertible snaks kept in order (somevalue as marker), bad property and value counted with the "reference " prefix, raw-kept dates counted, properties collected
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/wikigeul/internal/source"
)

func TestRefSnaks(t *testing.T) {
	rep := Report{Skipped: map[string]int{}, Raw: map[string]int{}}
	got := refSnaks(source.Reference{Snaks: []source.Qualifier{
		{Property: "P248", Type: "wikibase-entityid", Value: "Q36578", SnakType: "value"},
		{Property: "X1", Type: "string", Value: "a", SnakType: "value"},
		{Property: "P2", Type: "unknown", Value: "{}", SnakType: "value"},
		{Property: "P813", Type: "time", Value: "x", SnakType: "value"},
		{Property: "P50", SnakType: "somevalue"},
	}}, nil, &rep)
	if len(got) != 3 || got[0].Property != "P248" || got[1].Property != "P813" || !core.IsSlotMarker(got[2].Value) {
		t.Fatalf("%+v", got)
	}
	if rep.Skipped["reference 속성 ID: X1"] != 1 || rep.Skipped["reference 형식 미지원: unknown"] != 1 || rep.Raw["reference time: JSON 아님"] != 1 || rep.Markers["reference UNKNOWN-VALUE"] != 1 {
		t.Fatalf("%+v", rep)
	}
	if !slices.Equal(rep.Properties, []string{"P248", "Q36578", "P813", "P50"}) && !slices.Equal(rep.Properties, []string{"P248", "P813", "P50"}) {
		t.Fatalf("properties %v", rep.Properties)
	}
}
