//ff:func feature=port type=helper control=sequence
//ff:what wikimediaLanguage — the BCP 47 tag of a Wikimedia language code that is not one, one tag per code (docs/conversion.md 언어 코드 대응); others as given
package port

import "strings"

// wikimediaBCP47 maps the Wikimedia language codes Wikidata uses that are not
// valid BCP 47 tags to the tag they stand for (2026-10-26 full scan: these 14
// are every such code in labels, aliases, descriptions and monolingual text).
// Each tag rests on its IANA Language Subtag Registry entry; the private use
// subtags are GEUL definitions.
//
// The map is one-to-one over every code Wikidata accepts (term and
// monolingualtext languages, 2026-10-29): no two codes give the same tag, so
// wikimediaCode restores the code. When Wikidata also accepts the target tag
// itself as a code (sgs and bat-smg, sr-latn and sr-el, ...), the code that
// is not BCP 47 gets a private use subtag: the code without its hyphens
// (zh-classical, longer than 8 letters, as zhclass). eml is in no IANA entry
// (ISO 639-3 retired it in 2009 for egl and rgn), so it becomes the Romance
// collection roa with x-eml rather than Emilian egl.
var wikimediaBCP47 = map[string]string{
	"sr-el": "sr-Latn-x-srel", "bat-smg": "sgs-x-batsmg", "fiu-vro": "vro-x-fiuvro", "cbk-zam": "cbk-x-cbkzam",
	"simple": "en-simple", "eml": "roa-x-eml", "roa-rup": "rup-x-roarup", "zh-classical": "lzh-x-zhclass",
	"map-bms": "jv-x-bms", "roa-tara": "nap-x-tara",
	"de-formal": "de-x-formal", "es-formal": "es-x-formal", "hu-formal": "hu-x-formal", "nl-informal": "nl-x-informal",
}

// wikimediaLanguage returns the BCP 47 tag for the Wikimedia code lang
// (case-insensitive), or lang unchanged when it is not in wikimediaBCP47.
func wikimediaLanguage(lang string) string {
	if tag, ok := wikimediaBCP47[strings.ToLower(lang)]; ok {
		return tag
	}
	return lang
}
