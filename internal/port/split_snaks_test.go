//ff:func feature=port type=helper control=sequence
//ff:what TestSplitSnaks — source-describing properties (P248, P854, P1476 …) apart from citing ones (P813, P143, P227, unclassified), order kept
package port

import "testing"

func TestSplitSnaks(t *testing.T) {
	d, c := splitSnaks([]refSnak{{Property: "P813"}, {Property: "P854"}, {Property: "P143"}, {Property: "P248"}, {Property: "P227"}, {Property: "P1476"}, {Property: "P9999"}})
	if len(d) != 3 || d[0].Property != "P854" || d[1].Property != "P248" || d[2].Property != "P1476" ||
		len(c) != 4 || c[0].Property != "P813" || c[1].Property != "P143" || c[3].Property != "P9999" {
		t.Fatalf("%v %v", d, c)
	}
}
