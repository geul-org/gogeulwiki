//ff:func feature=cli type=helper control=sequence
//ff:what TestJobsOf — one job of the kind per number
package main

import "testing"

func TestJobsOf(t *testing.T) {
	js := jobsOf('Q', []uint64{3, 9})
	if len(js) != 2 || js[1] != (packJob{Kind: 'Q', Num: 9}) {
		t.Fatal(js)
	}
}
