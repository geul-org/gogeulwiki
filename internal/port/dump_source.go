//ff:func feature=port type=encoder control=sequence
//ff:what dumpSource — the HEAD SOURCE on the dump file: Origin a LOCAL bound to the dump URI, QUALIFIERs for the file's SHA-256 (P4092) and when it was received (P813)
package port

import "github.com/geul-org/gogeul/pkg/core"

// dumpSource writes the source of the whole document that the dump file is
// (README 출처 기록): a LOCAL bound to p.DumpURI by REFERENCE-NODE
// (localDefinition, declared in the BODY), SOURCE(#0, LOCAL) and on it
// QUALIFIER P4092 (checksum: the file's SHA-256, fileDigest; left out when
// unknown) and QUALIFIER P813 (retrieved: p.AcquiredAt, POSIX seconds). The
// file digest and the download time describe this run's copy, so they are
// citation QUALIFIERs, not facts of the source (edge/SOURCE 두 층). It
// returns the HEAD records and the BODY declaration; a malformed digest is an
// error.
func dumpSource(a IDs, p Provenance) (head, decl []core.Record, err error) {
	digest, err := fileDigest(p.DumpSHA256)
	if err != nil {
		return nil, nil, err
	}
	def, local := localDefinition(a, core.EncodeBytesValue(codecURIEven, []byte(p.DumpURI)))
	src, id := sourceEdge(a, 0, core.EncodeLocalRef(local))
	head = []core.Record{src}
	if !core.IsAbsent(digest) {
		p4092, _ := entityRef("P4092")
		head = append(head, qualifierEdge(a, p4092, id, digest))
	}
	p813, _ := entityRef("P813")
	head = append(head, qualifierEdge(a, p813, id, core.EncodePosixSeconds(p.AcquiredAt)))
	return head, []core.Record{def}, nil
}
