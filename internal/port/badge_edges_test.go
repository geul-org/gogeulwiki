//ff:func feature=port type=encoder control=sequence
//ff:what TestBadgeEdges — the badge predicate is bound once, one TRIPLE(badge, page, item) per item badge; a non-item badge is skipped; no badges give nothing
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestBadgeEdges(t *testing.T) {
	rep := Report{Skipped: map[string]int{}}
	ids := freshIDs(10)
	var pred uint64
	rs := badgeEdges(ids, 7, []string{"Q17437796", "P1", "x", "Q17437798"}, &pred, &rep)
	if len(rs) != 3 || pred != 11 || rep.Badges != 2 || rep.Skipped["sitelink badge: P1"] != 1 || rep.Skipped["sitelink badge: x"] != 1 ||
		rs[0].Codec != edgeReferenceNode || rs[1].Codec != edgeTriple || rs[2].Codec != edgeTriple {
		t.Fatalf("got %+v pred %d rep %+v", rs, pred, rep)
	}
	if len(rep.Notes) != 2 || rep.Notes[1] != (Note{Kind: "skipped", Reason: "sitelink badge: x", Property: "badge", Value: "x"}) {
		t.Fatalf("notes %+v", rep.Notes)
	}
	// TRIPLE #12: predicate LOCAL 11, subject the page LOCAL 7, object Q17437796, normal.
	if want := []core.Word{0x0088, 12, 0x0088, 11, 0x0088, 7, 0x0120, 0x010A, 0x1464, 0}; !slices.Equal(rs[1].Payload, want) {
		t.Fatalf("triple %04X want %04X", rs[1].Payload, want)
	}
	if rs := badgeEdges(ids, 8, []string{"Q1"}, &pred, &rep); len(rs) != 1 || rs[0].Codec != edgeTriple || pred != 11 {
		t.Fatalf("bound already: %+v", rs)
	}
	if rs := badgeEdges(ids, 8, nil, &pred, &rep); rs != nil {
		t.Fatal("no badges: want none")
	}
}
