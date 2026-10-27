//ff:func feature=port type=encoder control=sequence
//ff:what TestDumpSource — the dump is a REFERENCE-NODE LOCAL with a SOURCE on #0, QUALIFIER P4092 (when the digest is known) and P813; a bad digest fails
package port

import (
	"slices"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestDumpSource(t *testing.T) {
	p := prov()
	head, decl, err := dumpSource(freshIDs(1), p)
	if err != nil || len(head) != 2 || len(decl) != 1 || decl[0].Codec != edgeReferenceNode {
		t.Fatalf("no digest: %d %d %v", len(head), len(decl), err)
	}
	if !slices.Equal(head[0].Payload, []core.Word{0x0088, 3, 0x0088, 0, 0x0088, 2}) || !slices.Equal(head[1].Payload[:6], []core.Word{0x0088, 4, 0x00A8, 813, 0x0088, 3}) {
		t.Fatalf("SOURCE #3 #0 #2, P813 on #3: %04X %04X", head[0].Payload, head[1].Payload)
	}
	p.DumpSHA256 = strings.Repeat("ab", 32)
	head, _, err = dumpSource(freshIDs(1), p)
	if err != nil || len(head) != 3 || !slices.Equal(head[1].Payload[:6], []core.Word{0x0088, 4, 0x00A8, 4092, 0x0088, 3}) {
		t.Fatalf("digest: %d %v", len(head), err)
	}
	p.DumpSHA256 = "zz"
	if _, _, err := dumpSource(freshIDs(1), p); err == nil {
		t.Fatal("bad digest: want error")
	}
}
