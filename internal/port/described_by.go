//ff:func feature=port type=helper control=iteration dimension=1
//ff:what describedBy — for each LOCAL of a BODY, the snaks its TRIPLEs (P, LOCAL, value) and its URI binding (REFERENCE-NODE, read as P854) say of it
package port

import "github.com/geul-org/gogeul/pkg/core"

// describedBy collects what body says of each LOCAL node as reference snaks
// (describedSnak): a TRIPLE whose Subject is the LOCAL gives (Predicate,
// Object); a REFERENCE-NODE binding the LOCAL gives (P854, Object), the
// reference URL form of a URL-only source (urlOnly). It is the inverse of
// originNode.
func describedBy(body []core.Record) map[uint64][]string {
	out := map[uint64][]string{}
	for _, r := range body {
		if local, text, ok := describedSnak(r); ok {
			out[local] = append(out[local], text)
		}
	}
	return out
}
