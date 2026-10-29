//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what TestHeadRecords — SPEC-VERSION, DOC-ID, one TRANSFORMED-BY per tool, then the revision SOURCE and the dump SOURCE with its QUALIFIERs, all HEAD Edges on #0, their origins declared for the BODY; a bad Issuer, digest or commit fails
package port

import (
	"slices"
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"

	"github.com/geul-org/wikigeul/internal/source"
)

func TestHeadRecords(t *testing.T) {
	e := source.Entity{ID: "Q42", LastRevID: 7, Modified: 9}
	head, decl, err := headRecords([]core.Word{1, 0}, 42, e, prov(), freshIDs(1))
	if err != nil || len(head) != 8 || len(decl) != 3 {
		t.Fatalf("got %d %d %v", len(head), len(decl), err)
	}
	var codecs []uint8
	for _, r := range head {
		codecs = append(codecs, r.Codec)
	}
	if !slices.Equal(codecs, []uint8{core.MetaSpecVersion, core.MetaDocID, metaTransformedBy, metaTransformedBy, metaTransformedBy, edgeSource, edgeSource, edgeQualifier}) {
		t.Fatalf("codecs %02X", codecs)
	}
	if issuer, doc, err := core.DecodeDocID(head[1]); err != nil || !slices.Equal(issuer, []core.Word{1, 0}) || doc[3] != 42 {
		t.Fatalf("DOC-ID: %v %v %v", issuer, doc, err)
	}
	porter := append([]core.Word{0x0360, 0, 1, 0, 0, 1, 2, 0x0D62}, slices.Repeat([]core.Word{0x2222}, 10)...)
	if !slices.Equal(head[3].Payload, porter) {
		t.Fatalf("porter %04X", head[3].Payload)
	}
	for _, r := range head[5:] {
		if _, ok := core.HeadEdge(r, map[uint64]bool{4: true, 7: true}); !ok {
			t.Fatalf("not a HEAD Edge: %04X", r.Payload)
		}
	}
	if _, _, err := headRecords([]core.Word{0, 1}, 42, e, prov(), freshIDs(1)); err == nil {
		t.Fatal("non-canonical Issuer: want error")
	}
	withSum := prov()
	withSum.DumpSHA256 = "ab"
	if _, _, err := headRecords([]core.Word{1, 0}, 42, e, withSum, freshIDs(1)); err == nil {
		t.Fatal("short dump digest: want error")
	}
	bad := prov()
	bad.Tools[1].Commit = strings.Repeat("2", 7)
	if _, _, err := headRecords([]core.Word{1, 0}, 42, e, bad, freshIDs(1)); err == nil {
		t.Fatal("short commit: want error")
	}
}
