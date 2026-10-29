//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what TestSourceEdges — GND: the LOCAL's TRIPLEs (P248, P407), then SOURCE, then QUALIFIERs P227·P813; Britannica: SOURCE on the item with QUALIFIER P1417; expected snak sets noted
package port

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/wikigeul/internal/source"
)

func TestSourceEdges(t *testing.T) {
	rep := Report{Skipped: map[string]int{}, Raw: map[string]int{}}
	st := newSourcing()
	refs := q42Refs()
	rs := sourceEdges(freshIDs(10), 2, []source.Reference{refs[3], refs[0]}, st, nil, &rep)
	codecs := []uint8{}
	for _, r := range rs {
		codecs = append(codecs, r.Codec)
	}
	want := []uint8{edgeTriple, edgeTriple, edgeSource, edgeQualifier, edgeQualifier, edgeSource, edgeQualifier}
	if string(codecs) != string(want) || rs[0].Payload[3] != 248 || rs[1].Payload[3] != 407 || len(st.expect[2]) != 2 || rep.Sources != 2 || rep.SourceQuals != 3 {
		t.Fatalf("%v %+v", codecs, rep)
	}
	if o := rs[5].Payload[4:]; o[0] != 0x0120 || o[1] != core.Word(5375741>>16) || o[2] != core.Word(5375741&0xFFFF) {
		t.Fatalf("Britannica origin %04x", o)
	}
}
