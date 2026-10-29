//ff:func feature=cli type=command control=sequence
//ff:what TestBuildAll — documents of the jobs the dump has; an error stops it
package main

import "testing"

func TestBuildAll(t *testing.T) {
	db, schema := fixtureDB(t)
	c := packTestConfig(t, schema)
	bs, err := buildAll(db, c, packTestProv(t, c), []packJob{{Kind: 'Q', Num: 1}, {Kind: 'Q', Num: 999}, {Kind: 'Q', Num: 42}})
	if err != nil || len(bs) != 2 {
		t.Fatalf("%d %v", len(bs), err)
	}
	c.Schema = "1x"
	if _, err := buildAll(db, c, packTestProv(t, c), []packJob{{Kind: 'Q', Num: 1}}); err == nil {
		t.Fatal("bad schema: want error")
	}
}
