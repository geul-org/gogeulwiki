//ff:func feature=port type=helper control=iteration dimension=1
//ff:what wikimediaCode — the Wikimedia language code a BCP 47 tag came from: the inverse of wikimediaLanguage (docs/conversion.md 언어 코드 대응)
package port

import "strings"

// wikimediaCode returns the Wikimedia code whose tag (wikimediaBCP47) is tag,
// compared without case as BCP 47 tags are, or tag unchanged when no code maps
// to it: such a tag was the Wikimedia code itself. It undoes
// wikimediaLanguage for every code Wikidata accepts because that map is
// one-to-one and none of its tags is itself a Wikidata code.
func wikimediaCode(tag string) string {
	for code, t := range wikimediaBCP47 {
		if strings.EqualFold(t, tag) {
			return code
		}
	}
	return tag
}
