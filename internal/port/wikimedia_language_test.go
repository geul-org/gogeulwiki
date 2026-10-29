//ff:func feature=port type=helper control=iteration dimension=1
//ff:what TestWikimediaLanguage — mapped codes (standard, private use, any case) become their tags; BCP 47 codes and unknowns stay
package port

import "testing"

func TestWikimediaLanguage(t *testing.T) {
	for in, want := range map[string]string{
		"sr-el": "sr-Latn-x-srel", "simple": "en-simple", "zh-classical": "lzh-x-zhclass", "map-bms": "jv-x-bms",
		"DE-Formal": "de-x-formal", "eml": "roa-x-eml", "roa-rup": "rup-x-roarup", "bat-smg": "sgs-x-batsmg",
		"de": "de", "zh-Hant": "zh-Hant", "": "", "egl": "egl", "rup": "rup", "lzh": "lzh", "sgs": "sgs", "sr-latn": "sr-latn",
	} {
		if got := wikimediaLanguage(in); got != want {
			t.Fatalf("%q: %q want %q", in, got, want)
		}
	}
}
