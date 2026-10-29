//ff:func feature=port type=helper control=iteration dimension=1
//ff:what canonicalTag — a BCP 47 tag in canonical case (RFC 5646 2.1.1): script Title, region upper, the rest lower
package port

import "strings"

// canonicalTag writes tag in the canonical case of RFC 5646 2.1.1, which
// LANGUAGE-BCP47-CODE stores: lower case, except a 4-letter script subtag in
// title case and a 2-letter or 3-digit region subtag in upper case. The first
// subtag and everything after a singleton (extension, `x-` private use, or a
// tag that starts with one) stay lower case. Only case changes; the form is checked elsewhere (bcp47Form).
func canonicalTag(tag string) string {
	ps := strings.Split(strings.ToLower(tag), "-")
	for i := 1; i < len(ps) && len(ps[0]) > 1; i++ {
		p := ps[i]
		switch {
		case len(p) == 1:
			return strings.Join(ps, "-")
		case len(p) == 4 && p[0] >= 'a' && p[0] <= 'z' && strings.Trim(p, "abcdefghijklmnopqrstuvwxyz") == "":
			ps[i] = strings.ToUpper(p[:1]) + p[1:]
		case len(p) == 2 && strings.Trim(p, "abcdefghijklmnopqrstuvwxyz") == "" || len(p) == 3 && strings.Trim(p, "0123456789") == "":
			ps[i] = strings.ToUpper(p)
		}
	}
	return strings.Join(ps, "-")
}
