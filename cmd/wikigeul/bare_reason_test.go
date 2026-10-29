//ff:func feature=cli type=helper control=sequence
//ff:what TestBareReason — qualifier and reference prefixes are dropped, a main-value reason stays
package main

import "testing"

func TestBareReason(t *testing.T) {
	if bareReason("qualifier 개체: L") != "개체: L" || bareReason("reference time: 날짜 형식") != "time: 날짜 형식" || bareReason("개체: E") != "개체: E" {
		t.Fatal("prefix")
	}
}
