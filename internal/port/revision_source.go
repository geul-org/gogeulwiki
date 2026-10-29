//ff:func feature=port type=encoder control=sequence
//ff:what revisionSource — the HEAD SOURCE on the entity's revision: Origin a LOCAL described by publication date (P577) and permanent URL (P854); nothing when the dump gives no revision
package port

import (
	"github.com/geul-org/gogeul/pkg/core"

	"github.com/geul-org/wikigeul/internal/source"
)

// revisionSource writes the source of the whole document that the entity e
// is (README 출처 기록): the revision the dump holds, a LOCAL described like
// a reference's source (originNode rule 3) by TRIPLE P577 (publication date:
// when the revision was saved, POSIX seconds) and TRIPLE P854 (its permanent
// URL, revisionURI) — a URL-only description when the save time is unknown
// is a REFERENCE-NODE (localDefinition) — and SOURCE(#0, LOCAL) on it. It
// returns the HEAD records and the BODY records declaring the LOCAL; both
// are empty when e has no revision ID.
func revisionSource(a IDs, e source.Entity) (head, decl []core.Record) {
	if e.LastRevID == 0 {
		return nil, nil
	}
	p577, _ := entityRef("P577")
	p854, _ := entityRef("P854")
	url := refSnak{Property: "P854", Prop: p854, Value: core.EncodeBytesValue(codecURIEven, []byte(revisionURI(e.LastRevID)))}
	describe := []refSnak{url}
	if e.Modified != 0 {
		describe = []refSnak{{Property: "P577", Prop: p577, Value: core.EncodePosixSeconds(e.Modified)}, url}
	}
	var local uint64
	if urlOnly(describe) {
		var def core.Record
		def, local = localDefinition(a, url.Value)
		decl = []core.Record{def}
	} else {
		decl, local = describedNode(a, describe)
	}
	src, _ := sourceEdge(a, 0, core.EncodeLocalRef(local))
	return []core.Record{src}, decl
}
