//ff:func feature=cli type=command control=sequence
//ff:what TestCheckFacts — a converted document and its list have no Facts; odd bytes give an Error on the record named
package main

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/resolve"
	"github.com/geul-org/gogeul/pkg/schema"
)

func TestCheckFacts(t *testing.T) {
	db, schemaName := fixtureDB(t)
	c := packTestConfig(t, schemaName)
	b, _, _, err := buildOne(db, c, packTestProv(t, c), packJob{Kind: 'Q', Num: 42})
	list, _ := resolve.FirstVersionList(b, c.RecordedAt)
	if fs := checkFacts(b, list); err != nil || len(fs) != 0 {
		t.Fatalf("%v %+v", err, fs)
	}
	fs := checkFacts(b, []byte{1})
	if len(fs) == 0 || fs[0].Doc != "list" || fs[0].Severity != schema.Error {
		t.Fatalf("%+v", fs)
	}
}
