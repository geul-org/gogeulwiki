//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what TestJobAfter — anything after nothing; Q after P; by number within a kind
package main

import "testing"

func TestJobAfter(t *testing.T) {
	for _, c := range []struct {
		k  byte
		n  uint64
		ak byte
		an uint64
		w  bool
	}{{'P', 1, 0, 0, true}, {'Q', 1, 'P', 9, true}, {'P', 9, 'Q', 1, false}, {'Q', 5, 'Q', 4, true}, {'Q', 4, 'Q', 4, false}} {
		if jobAfter(c.k, c.n, c.ak, c.an) != c.w {
			t.Errorf("%c%d after %c%d", c.k, c.n, c.ak, c.an)
		}
	}
}
