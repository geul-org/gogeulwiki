//ff:func feature=port type=helper control=sequence
//ff:what TestSnakSet — order does not count, duplicates do, the input is not reordered
package port

import "testing"

func TestSnakSet(t *testing.T) {
	in := []string{"b", "a"}
	if snakSet(in) != snakSet([]string{"a", "b"}) || snakSet(in) != "a;b;" || in[0] != "b" {
		t.Fatalf("%q %v", snakSet(in), in)
	}
	if snakSet([]string{"a", "a"}) == snakSet([]string{"a"}) {
		t.Fatal("duplicates must count")
	}
}
