//ff:func feature=cli type=helper control=sequence
//ff:what TestSendAll — every job in order
package main

import "testing"

func TestSendAll(t *testing.T) {
	var got []uint64
	sendAll([]packJob{{Num: 1}, {Num: 2}}, func(j packJob) { got = append(got, j.Num) })
	if len(got) != 2 || got[1] != 2 {
		t.Fatal(got)
	}
}
