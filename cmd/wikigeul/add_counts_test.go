//ff:func feature=cli type=helper control=sequence
//ff:what TestAddCounts — counts add up per reason
package main

import "testing"

func TestAddCounts(t *testing.T) {
	sum := map[string]int{"a": 1}
	addCounts(sum, map[string]int{"a": 2, "b": 1})
	addCounts(sum, nil)
	if sum["a"] != 3 || sum["b"] != 1 || len(sum) != 2 {
		t.Fatalf("%v", sum)
	}
}
