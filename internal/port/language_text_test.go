//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what TestLanguageText — a Wikimedia code is mapped; a tag-shaped code gives the docs' "Paris"@fr words; with a registry a numbered tag is a LANGUAGE-SHORT-CODE and an unnumbered one stays BCP 47; malformed codes are false
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestLanguageText(t *testing.T) {
	r, ok := languageText(nil, "fr", "Paris")
	ws, _ := core.EncodeRecord(r)
	if !ok || !slices.Equal(ws, []core.Word{0x0ACB, 0x08C6, 0x6672, 0x5061, 0x7269, 0x7300}) {
		t.Fatalf("got %04X %v", ws, ok)
	}
	r, _ = languageText(nil, "sr-el", "Beograd")
	lang, _, err := core.DecodeLanguageText(r)
	if b, _ := core.DecodeBytesValue(lang, core.CodecLanguageBCP47Even); err != nil || string(b) != "sr-Latn-x-srel" {
		t.Fatalf("sr-el: %q %v", b, err)
	}
	if _, ok := languageText(nil, "zh-Hant", "x"); !ok {
		t.Fatal("zh-Hant: want ok")
	}
	langs := Languages{"en-simple": 0x40E8, "fr": 0x0597}
	r, _ = languageText(langs, "simple", "Paris")
	if ws, _ := core.EncodeRecord(r); !slices.Equal(ws[:3], []core.Word{0x0ACB, 0x08C8, 0x40E8}) {
		t.Fatalf("simple → en-simple in the registry: %04X", ws)
	}
	r, _ = languageText(langs, "de-formal", "Sie")
	lang, _, _ = core.DecodeLanguageText(r)
	if b, _ := core.DecodeBytesValue(lang, core.CodecLanguageBCP47Even); string(b) != "de-x-formal" {
		t.Fatalf("de-formal → de-x-formal stays BCP 47: %q", b)
	}
	for _, bad := range []string{"", "e", "en_US", "en-", "toolongcode", "ko KR"} {
		if _, ok := languageText(nil, bad, "x"); ok {
			t.Fatalf("%q: want false", bad)
		}
	}
}
