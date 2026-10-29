//ff:func feature=port type=helper control=iteration dimension=1
//ff:what TestWikimediaLanguageInjective — over every code Wikidata accepts, no two codes share a tag and wikimediaCode gives each code back
package port

import (
	"os"
	"strings"
	"testing"
)

func TestWikimediaLanguageInjective(t *testing.T) {
	b, err := os.ReadFile("../../testdata/wikidata-languages.txt")
	if err != nil {
		t.Fatal(err)
	}
	seen := map[string]string{}
	for _, code := range strings.Split(strings.TrimSpace(string(b)), "\n") {
		if strings.HasPrefix(code, "#") {
			continue
		}
		tag := wikimediaLanguage(code)
		if other, ok := seen[strings.ToLower(tag)]; ok {
			t.Fatalf("%q and %q: same tag %q", other, code, tag)
		}
		seen[strings.ToLower(tag)] = code
		if back := wikimediaCode(tag); back != code {
			t.Fatalf("%q -> %q -> %q", code, tag, back)
		}
	}
	if len(seen) < 1400 {
		t.Fatalf("only %d codes read", len(seen))
	}
	for code, tag := range wikimediaBCP47 {
		if seen[strings.ToLower(tag)] != code {
			t.Fatalf("%q: not in the Wikidata code list", code)
		}
	}
}
