//ff:func feature=cli type=command control=sequence
//ff:what TestConvertOne — a document with its list, both encoded, and its document number; a missing entity and a bad schema carry their state; strict checks find nothing on a good document; a panic becomes the error
package main

import (
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/resolve"
)

func TestConvertOne(t *testing.T) {
	db, schema := fixtureDB(t)
	c := packTestConfig(t, schema)
	s := packTestStream(t)
	r := convertOne(db, c, packTestProv(t, c), s, packJob{Kind: 'Q', Num: 42, Seq: 3})
	list, _ := resolve.FirstVersionList(r.Raw, c.RecordedAt)
	if r.Err != nil || r.Missing || r.Doc != 0x510000000000002a || string(r.List.Bytes) != string(list) || r.Version.Raw != len(r.Raw) || !r.Version.HashOK || r.Job.Seq != 3 {
		t.Fatalf("%+v", r.Err)
	}
	if r := convertOne(db, c, packTestProv(t, c), s, packJob{Kind: 'Q', Num: 999}); !r.Missing || r.Err != nil {
		t.Fatal("missing")
	}
	strict := c
	strict.Strict = true
	if r := convertOne(db, strict, packTestProv(t, c), s, packJob{Kind: 'Q', Num: 42}); r.Err != nil || r.Facts != nil && len(r.Facts) != 0 {
		t.Fatalf("strict: %v %+v", r.Err, r.Facts)
	}
	if r := convertOne(nil, c, packTestProv(t, c), s, packJob{Kind: 'Q', Num: 42}); r.Err == nil || !strings.HasPrefix(r.Err.Error(), "panic: ") {
		t.Fatalf("nil database: want a recovered panic, got %v", r.Err)
	}
	c.Schema = "1x"
	if r := convertOne(db, c, packTestProv(t, c), s, packJob{Kind: 'Q', Num: 42}); r.Err == nil {
		t.Fatal("bad schema: want error")
	}
}
