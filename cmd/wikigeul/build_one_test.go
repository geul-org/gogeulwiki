//ff:func feature=cli type=command control=sequence
//ff:what TestBuildOne — an entity of the dump gives its document; a missing one is false without error; a bad schema is an error
package main

import "testing"

func TestBuildOne(t *testing.T) {
	db, schema := fixtureDB(t)
	c := packTestConfig(t, schema)
	b, rep, ok, err := buildOne(db, c, packTestProv(t, c), packJob{Kind: 'Q', Num: 42})
	if err != nil || !ok || len(b) == 0 || rep.ID != "Q42" || rep.Labels == 0 {
		t.Fatalf("%v %v", rep, err)
	}
	if _, rep, ok, err := buildOne(db, c, packTestProv(t, c), packJob{Kind: 'Q', Num: 999}); ok || err != nil || !rep.Missing {
		t.Fatal("missing")
	}
	c.Schema = "1x"
	if _, _, ok, err := buildOne(db, c, packTestProv(t, c), packJob{Kind: 'Q', Num: 42}); ok || err == nil {
		t.Fatal("bad schema: want error")
	}
}
