//ff:func feature=port type=helper control=iteration dimension=1
//ff:what TestWikimediaCode — every mapped tag (any case) gives back its code; a tag no code maps to is the code itself
package port

import "testing"

func TestWikimediaCode(t *testing.T) {
	for in, want := range map[string]string{
		"sr-Latn-x-srel": "sr-el", "SR-LATN-X-SREL": "sr-el", "sgs-x-batsmg": "bat-smg", "vro-x-fiuvro": "fiu-vro",
		"cbk-x-cbkzam": "cbk-zam", "en-simple": "simple", "roa-x-eml": "eml", "rup-x-roarup": "roa-rup",
		"lzh-x-zhclass": "zh-classical", "jv-x-bms": "map-bms", "nap-x-tara": "roa-tara", "de-x-formal": "de-formal",
		"es-x-formal": "es-formal", "hu-x-formal": "hu-formal", "nl-x-informal": "nl-informal",
		"sgs": "sgs", "sr-latn": "sr-latn", "egl": "egl", "lzh": "lzh", "de": "de", "": "",
	} {
		if got := wikimediaCode(in); got != want {
			t.Fatalf("%q: %q want %q", in, got, want)
		}
	}
}
