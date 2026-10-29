//ff:func feature=cli type=helper control=sequence
//ff:what TestPageKind — every entity of a kind in the bound, in order; a bad schema fails
package main

import (
	"testing"
)

func TestPageKind(t *testing.T) {
	db, schema := fixtureDB(t)
	c := config{Schema: schema, Dump: 1}
	var got []uint64
	if err := pageKind(db, c, 'Q', 2, 42, func(j packJob) { got = append(got, j.Num) }); err != nil || len(got) != 4 || got[0] != 3 || got[3] != 42 {
		t.Fatalf("%v %v", got, err)
	}
	if err := pageKind(db, config{Schema: "1x", Dump: 1}, 'Q', 0, 9, func(packJob) {}); err == nil {
		t.Fatal("bad schema: want error")
	}
}
