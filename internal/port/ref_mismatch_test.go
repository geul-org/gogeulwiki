//ff:func feature=port type=helper control=sequence
//ff:what TestRefMismatch — equal lists in any order match; a changed set, a missing Edge and an extra Edge each count
package port

import "testing"

func TestRefMismatch(t *testing.T) {
	if n := refMismatch(map[uint64][]string{1: {"a", "b"}}, map[uint64][]string{1: {"b", "a"}}); n != 0 {
		t.Fatalf("equal: %d", n)
	}
	if n := refMismatch(map[uint64][]string{1: {"a"}, 2: {"c"}}, map[uint64][]string{1: {"x"}, 3: {"c"}}); n != 3 {
		t.Fatalf("differ: %d", n)
	}
}
