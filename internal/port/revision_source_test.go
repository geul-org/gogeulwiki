//ff:func feature=port type=encoder control=sequence
//ff:what TestRevisionSource — a revision with its time is a LOCAL with TRIPLE P577 and P854 and a SOURCE on #0; without the time a REFERENCE-NODE; without a revision nothing
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"

	"github.com/geul-org/gogeulwiki/internal/source"
)

func TestRevisionSource(t *testing.T) {
	head, decl := revisionSource(freshIDs(1), source.Entity{LastRevID: 7, Modified: 9})
	if len(head) != 1 || len(decl) != 2 || decl[0].Codec != edgeTriple || decl[1].Codec != edgeTriple {
		t.Fatalf("described: %d %d", len(head), len(decl))
	}
	if !slices.Equal(head[0].Payload, []core.Word{0x0088, 4, 0x0088, 0, 0x0088, 1}) {
		t.Fatalf("SOURCE #4 #0 #1: %04X", head[0].Payload)
	}
	if !slices.Equal(decl[0].Payload[2:6], []core.Word{0x00A8, 577, 0x0088, 1}) || !slices.Equal(decl[1].Payload[2:6], []core.Word{0x00A8, 854, 0x0088, 1}) {
		t.Fatalf("P577 then P854 on #1: %04X %04X", decl[0].Payload, decl[1].Payload)
	}
	head, decl = revisionSource(freshIDs(1), source.Entity{LastRevID: 7})
	if len(head) != 1 || len(decl) != 1 || decl[0].Codec != edgeReferenceNode {
		t.Fatalf("URL only: %d %d", len(head), len(decl))
	}
	if head, decl := revisionSource(freshIDs(1), source.Entity{}); head != nil || decl != nil {
		t.Fatal("no revision: want nothing")
	}
}
