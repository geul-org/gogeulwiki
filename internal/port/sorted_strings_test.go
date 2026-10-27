//ff:func feature=port type=helper control=sequence
//ff:what TestSortedStrings — a sorted copy, the input untouched
package port

import (
	"slices"
	"testing"
)

func TestSortedStrings(t *testing.T) {
	in := []string{"b", "a"}
	if got := sortedStrings(in); !slices.Equal(got, []string{"a", "b"}) || in[0] != "b" {
		t.Fatal(got)
	}
}
