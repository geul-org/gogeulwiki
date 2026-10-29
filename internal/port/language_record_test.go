//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what TestLanguageRecord — a tag the registry has (any case) is its LANGUAGE-SHORT-CODE; one it lacks, and every tag without a registry, is a LANGUAGE-BCP47-CODE in canonical case
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestLanguageRecord(t *testing.T) {
	langs := Languages{"en": 0x0784, "pt-br": 0x4130, "en-simple": 0x40e8}
	for tag, code := range map[string]core.Word{"en": 0x0784, "pt-br": 0x4130, "PT-BR": 0x4130, "pt-BR": 0x4130, "en-simple": 0x40E8} {
		ws, _ := core.EncodeRecord(languageRecord(langs, tag))
		if !slices.Equal(ws, []core.Word{0x08C8, code}) {
			t.Errorf("%s: %04X", tag, ws)
		}
	}
	for tag, want := range map[string]string{"de-x-formal": "de-x-formal", "en-uk": "en-UK", "nl-u-sd-bebru": "nl-u-sd-bebru", "sr-latn-x-srel": "sr-Latn-x-srel"} {
		b1, err1 := core.DecodeBytesValue(languageRecord(langs, tag), core.CodecLanguageBCP47Even)
		b2, err2 := core.DecodeBytesValue(languageRecord(nil, tag), core.CodecLanguageBCP47Even)
		if err1 != nil || err2 != nil || string(b1) != want || string(b2) != want {
			t.Errorf("%s: %q %q %v %v", tag, b1, b2, err1, err2)
		}
	}
	if r := languageRecord(nil, "en"); r.Codec&^1 != core.CodecLanguageBCP47Even {
		t.Errorf("no registry: %+v", r)
	}
}
