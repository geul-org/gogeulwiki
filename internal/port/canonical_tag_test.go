//ff:func feature=port type=helper control=iteration dimension=1
//ff:what TestCanonicalTag — script Title, region upper, the rest lower; nothing after a singleton or in the first subtag changes case
package port

import "testing"

func TestCanonicalTag(t *testing.T) {
	for in, want := range map[string]string{
		"en": "en", "EN": "en", "pt-br": "pt-BR", "zh-hant": "zh-Hant", "ZH-HANT-tw": "zh-Hant-TW",
		"es-419": "es-419", "sr-latn-x-srel": "sr-Latn-x-srel", "de-X-FORMAL": "de-x-formal",
		"nl-u-sd-bebru": "nl-u-sd-bebru", "en-uk": "en-UK", "cmn-latn-tw-wadegile": "cmn-Latn-TW-wadegile",
		"be-tarask": "be-tarask", "i-klingon": "i-klingon", "en-1994": "en-1994",
		"x-abcd": "x-abcd", "ab-x1y2": "ab-x1y2",
	} {
		if got := canonicalTag(in); got != want {
			t.Errorf("%s: got %s, want %s", in, got, want)
		}
	}
}
