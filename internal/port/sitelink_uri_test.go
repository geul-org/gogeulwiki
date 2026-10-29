//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what TestSitelinkURI — the canonical page URL per site (spaces as _, MediaWiki encoding, exception hosts); an unknown site is false
package port

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"

	"github.com/geul-org/wikigeul/internal/source"
)

func TestSitelinkURI(t *testing.T) {
	for want, s := range map[string]source.Sitelink{
		"https://en.wikipedia.org/wiki/Douglas_Adams":                        {Site: "enwiki", Title: "Douglas Adams"},
		"https://commons.wikimedia.org/wiki/Category:Douglas_Adams":          {Site: "commonswiki", Title: "Category:Douglas Adams"},
		"https://en.wikiquote.org/wiki/What%3F/A_%26_B%27s_100%25":           {Site: "enwikiquote", Title: "What?/A & B's 100%"},
		"https://ko.wikipedia.org/wiki/%EB%8D%94%EA%B8%80%EB%9F%AC%EC%8A%A4": {Site: "kowiki", Title: "더글러스"},
		"https://zh-min-nan.wikipedia.org/wiki/AC/DC_(band)":                 {Site: "zh_min_nanwiki", Title: "AC/DC (band)"},
	} {
		r, ok := sitelinkURI(s)
		if b, err := core.DecodeBytesValue(r, codecURIEven); !ok || err != nil || string(b) != want {
			t.Fatalf("%+v: got %q %v want %q", s, b, ok, want)
		}
	}
	if _, ok := sitelinkURI(source.Sitelink{Site: "xxsite", Title: "x"}); ok {
		t.Fatal("unknown site: want false")
	}
}
