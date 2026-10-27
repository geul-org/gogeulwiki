//ff:func feature=port type=helper control=sequence
//ff:what urlOnly — whether a source is described by one reference URL (P854) written as URI and nothing else
package port

import "github.com/geul-org/gogeul/pkg/core"

// urlOnly reports whether describe is exactly one P854 (reference URL) snak
// whose value is a URI Value (url values without a scheme stay UTF8-TEXT).
// Such a source is a LOCAL bound to the URI by REFERENCE-NODE, like a
// sitelink page (README 출처).
func urlOnly(describe []refSnak) bool {
	if len(describe) != 1 || describe[0].Property != "P854" {
		return false
	}
	v := describe[0].Value
	return v.Kind == core.KindValue && len(v.Issuer) == 0 && v.Codec&^1 == codecURIEven
}
