//ff:func feature=port type=helper control=iteration dimension=1
//ff:what wikiSiteHost — the host of a Wikimedia site from its global ID: the listed exceptions, else language code + project family; false when neither fits
package port

import (
	"regexp"
	"strings"
)

// wikiSiteHosts are the site IDs whose host the family rule does not give
// (Wikimedia SiteMatrix, meta.wikimedia.org action=sitematrix, 2026-10-27:
// every other site ID in the 20261019 dump's sitelinks follows the rule).
var wikiSiteHosts = map[string]string{
	"be_x_oldwiki":      "be-tarask.wikipedia.org",
	"commonswiki":       "commons.wikimedia.org",
	"foundationwiki":    "foundation.wikimedia.org",
	"mediawikiwiki":     "www.mediawiki.org",
	"metawiki":          "meta.wikimedia.org",
	"outreachwiki":      "outreach.wikimedia.org",
	"sourceswiki":       "wikisource.org",
	"specieswiki":       "species.wikimedia.org",
	"wikidatawiki":      "www.wikidata.org",
	"wikifunctionswiki": "www.wikifunctions.org",
	"wikimaniawiki":     "wikimania.wikimedia.org",
}

// wikiFamilies pairs each site ID suffix with its project domain; "wiki"
// (Wikipedia) is last so that the longer suffixes are tried first.
var wikiFamilies = [][2]string{
	{"wiktionary", "wiktionary.org"}, {"wikibooks", "wikibooks.org"}, {"wikinews", "wikinews.org"},
	{"wikiquote", "wikiquote.org"}, {"wikisource", "wikisource.org"}, {"wikiversity", "wikiversity.org"},
	{"wikivoyage", "wikivoyage.org"}, {"wiki", "wikipedia.org"},
}

// wikiSiteLang is the language part of a site ID (zh_min_nan, simple, en).
var wikiSiteLang = regexp.MustCompile(`^[a-z][a-z0-9_]*$`)

// wikiSiteHost returns the host of the Wikimedia site with global ID site
// (sitelink "site"): the exception in wikiSiteHosts, else the language part
// with _ as - under the project family of the suffix (enwiki →
// en.wikipedia.org, zh_min_nanwikisource → zh-min-nan.wikisource.org). It is
// false for an ID that fits neither.
func wikiSiteHost(site string) (string, bool) {
	if host, ok := wikiSiteHosts[site]; ok {
		return host, true
	}
	for _, f := range wikiFamilies {
		lang, ok := strings.CutSuffix(site, f[0])
		if ok && wikiSiteLang.MatchString(lang) {
			return strings.ReplaceAll(lang, "_", "-") + "." + f[1], true
		}
	}
	return "", false
}
