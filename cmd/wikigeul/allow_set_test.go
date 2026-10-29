//ff:func feature=cli type=helper control=sequence
//ff:what TestAllowSet — every entry is in the set, nothing else
package main

import "testing"

func TestAllowSet(t *testing.T) {
	s := allowSet([]string{"raw:time: 0년", "warn:E7"})
	if len(s) != 2 || !s["raw:time: 0년"] || !s["warn:E7"] || s["warn:E6"] {
		t.Fatalf("%v", s)
	}
}
