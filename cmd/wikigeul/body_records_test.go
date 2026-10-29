//ff:func feature=cli type=helper control=sequence
//ff:what TestBodyRecords — the BODY entries' records in order; none for an empty BODY
package main

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

func TestBodyRecords(t *testing.T) {
	a := core.EncodeLocalRef(1)
	b := core.EncodeLocalRef(2)
	got := bodyRecords(inspect.Document{Head: []inspect.Entry{{Record: b}}, Body: []inspect.Entry{{Record: a}, {Record: b}}})
	if len(got) != 2 || got[0].Payload[0] != 1 || got[1].Payload[0] != 2 {
		t.Fatalf("%+v", got)
	}
	if got := bodyRecords(inspect.Document{}); len(got) != 0 {
		t.Fatalf("empty: %+v", got)
	}
}
