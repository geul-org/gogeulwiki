//ff:func feature=cli type=command control=sequence
//ff:what TestDictSamples — the first Samples Q entities of the range or of the ID list; a bad schema fails
package main

import "testing"

func TestDictSamples(t *testing.T) {
	db, schema := fixtureDB(t)
	c := packTestConfig(t, schema)
	c.Samples = 3
	if bs, err := dictSamples(db, c, packTestProv(t, c), nil); err != nil || len(bs) != 3 {
		t.Fatalf("%d %v", len(bs), err)
	}
	ids := []packJob{{Kind: 'P', Num: 31}, {Kind: 'Q', Num: 42}}
	if bs, err := dictSamples(db, c, packTestProv(t, c), ids); err != nil || len(bs) != 1 {
		t.Fatalf("ID list: %d %v", len(bs), err)
	}
	c.Schema = "1x"
	if _, err := dictSamples(db, c, packTestProv(t, c), nil); err == nil {
		t.Fatal("bad schema: want error")
	}
}
