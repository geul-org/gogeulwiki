//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what headRecords — the HEAD of a ported document: SPEC-VERSION, DOC-ID, one TRANSFORMED-BY per tool, the SOURCEs on the whole document (entity revision, dump file) with their QUALIFIERs; and the BODY records declaring their origins
package port

import (
	"github.com/geul-org/gogeul/pkg/core"

	"github.com/geul-org/gogeulwiki/internal/source"
)

// headRecords writes the HEAD of the document of entity e numbered num
// under issuer: SPEC-VERSION, DOC-ID DOC-POINTER(issuer, num, 0), one
// TRANSFORMED-BY per tool in p.Tools (the tool's Reference, toolRef, and its
// commit, CommitDigest), then the HEAD Edges (document HEAD): the SOURCE on
// the entity's revision (revisionSource) and the SOURCE on the dump file with
// its QUALIFIERs (dumpSource). It also returns the BODY records that declare
// those SOURCEs' origin LOCALs, to open the BODY with. Edge Ids come from a
// in that order. A malformed Issuer, commit or digest is an error. No
// wall-clock time is written, so the same input and tools give the same
// document.
func headRecords(issuer []core.Word, num uint64, e source.Entity, p Provenance, a IDs) (head, decl []core.Record, err error) {
	id, err := core.EncodeDocID(issuer, []core.Word{core.Word(num >> 48), core.Word(num >> 32), core.Word(num >> 16), core.Word(num)})
	if err != nil {
		return nil, nil, err
	}
	head = []core.Record{core.EncodeSpecVersion(core.SpecVersion), id}
	for _, t := range p.Tools {
		commit, err := CommitDigest(t.Commit)
		if err != nil {
			return nil, nil, err
		}
		head = append(head, compose(core.KindMeta, metaTransformedBy, toolRef(t.Doc), commit))
	}
	rev, revDecl := revisionSource(a, e)
	dump, dumpDecl, err := dumpSource(a, p)
	if err != nil {
		return nil, nil, err
	}
	return append(append(head, rev...), dump...), append(revDecl, dumpDecl...), nil
}
