//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what TestMonolingualRecord — text and language become UTF8-LANGUAGE-TEXT (a leading U+FEFF kept); bad JSON and a malformed language give reasons
package port

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestMonolingualRecord(t *testing.T) {
	r, reason := monolingualRecord(`{"text":"Bundesrepublik Deutschland","language":"de"}`, nil)
	lang, text, err := core.DecodeLanguageText(r)
	if reason != "" || err != nil || string(text) != "Bundesrepublik Deutschland" || lang.Payload[0] != 0x6465 {
		t.Fatalf("%q %v", reason, err)
	}
	r, reason = monolingualRecord(`{"text":"\ufeffAbecásia (GE)","language":"pt"}`, nil)
	if _, text, err := core.DecodeLanguageText(r); reason != "" || err != nil || string(text) != "\ufeffAbecásia (GE)" {
		t.Fatalf("leading U+FEFF: %q %q %v", reason, text, err)
	}
	for want, v := range map[string]string{
		"monolingualtext: JSON 아님":   "{",
		"monolingualtext: 언어 코드 x y": `{"text":"a","language":"x y"}`,
	} {
		if _, got := monolingualRecord(v, nil); got != want {
			t.Fatalf("%q: %q", v, got)
		}
	}
}
