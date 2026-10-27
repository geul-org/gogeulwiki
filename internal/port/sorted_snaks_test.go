//ff:func feature=port type=helper control=sequence
//ff:what TestSortedSnaks — by property number, one property's values keep their order, the input is not reordered
package port

import "testing"

func TestSortedSnaks(t *testing.T) {
	a, _ := entityRef("Q1")
	b, _ := entityRef("Q2")
	in := []refSnak{{Property: "P2093", Value: a}, {Property: "P854"}, {Property: "P2093", Value: b}, {Property: "P50"}}
	got := sortedSnaks(in)
	if got[0].Property != "P50" || got[1].Property != "P854" || got[2].Value.Payload[0] != 1 || got[3].Value.Payload[0] != 2 || in[0].Property != "P2093" {
		t.Fatalf("%v", got)
	}
}
