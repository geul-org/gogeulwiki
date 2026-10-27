//ff:func feature=port type=helper control=iteration dimension=1
//ff:what TestWikiSiteHost — family rule per project with _ as -, the SiteMatrix exceptions; unknown suffixes and bad language parts are false
package port

import "testing"

func TestWikiSiteHost(t *testing.T) {
	for site, want := range map[string]string{
		"enwiki": "en.wikipedia.org", "kowiktionary": "ko.wiktionary.org", "dewikibooks": "de.wikibooks.org",
		"frwikinews": "fr.wikinews.org", "arwikiquote": "ar.wikiquote.org", "zh_min_nanwikisource": "zh-min-nan.wikisource.org",
		"enwikiversity": "en.wikiversity.org", "jawikivoyage": "ja.wikivoyage.org", "simplewiki": "simple.wikipedia.org",
		"be_x_oldwiki": "be-tarask.wikipedia.org", "commonswiki": "commons.wikimedia.org", "wikidatawiki": "www.wikidata.org",
		"sourceswiki": "wikisource.org", "mediawikiwiki": "www.mediawiki.org", "specieswiki": "species.wikimedia.org",
	} {
		if got, ok := wikiSiteHost(site); !ok || got != want {
			t.Fatalf("%s: %q %v want %q", site, got, ok, want)
		}
	}
	for _, site := range []string{"xxsite", "wiki", "EnWiki", "_wiki", "en-wiki", ""} {
		if got, ok := wikiSiteHost(site); ok {
			t.Fatalf("%q: got %q, want false", site, got)
		}
	}
}
