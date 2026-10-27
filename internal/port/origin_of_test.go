//ff:func feature=port type=encoder control=sequence
//ff:what TestOriginOf — nothing describing is ABSENT, a lone P248 item is that item, a P248 property value or more snaks make a LOCAL node
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestOriginOf(t *testing.T) {
	rep := Report{}
	st := newSourcing()
	a := freshIDs(1)
	if o, decl := originOf(a, nil, st, &rep); !core.IsAbsent(o) || decl != nil || rep.OriginNone != 1 {
		t.Fatalf("none: %+v %v", o, decl)
	}
	p248, _ := entityRef("P248")
	q, _ := entityRef("Q36578")
	if o, decl := originOf(a, []refSnak{{Property: "P248", Prop: p248, Value: q}}, st, &rep); !slices.Equal(wordsOf(o), wordsOf(q)) || decl != nil || rep.OriginQIDs != 1 {
		t.Fatalf("item: %+v", o)
	}
	pid, _ := entityRef("P31")
	o, decl := originOf(a, []refSnak{{Property: "P248", Prop: p248, Value: pid}}, st, &rep)
	if id, err := core.DecodeLocalRef(o); err != nil || id != 1 || len(decl) != 1 || rep.OriginLocals != 1 {
		t.Fatalf("property value: %+v %v", o, decl)
	}
}
