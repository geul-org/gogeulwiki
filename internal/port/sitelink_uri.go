//ff:func feature=port type=encoder control=sequence
//ff:what sitelinkURI — a sitelink's page URL as a URI Value: https:// + site host + /wiki/ + the title with _ for spaces, encoded as MediaWiki does; false for an unknown site
package port

import (
	"strings"

	"github.com/geul-org/gogeul/pkg/core"

	"github.com/geul-org/gogeulwiki/internal/source"
)

// sitelinkURI builds the canonical page URL of s (README sitelink): the
// site's host (wikiSiteHost), the article path /wiki/ every Wikimedia site
// uses and the title with spaces as _ percent-encoded the way MediaWiki
// links it (isTitleByte, upper-case hex). It is false when the site ID has no
// known host.
func sitelinkURI(s source.Sitelink) (core.Record, bool) {
	host, ok := wikiSiteHost(s.Site)
	if !ok {
		return core.Record{}, false
	}
	title := percentEncode(strings.ReplaceAll(s.Title, " ", "_"), isTitleByte)
	return core.EncodeBytesValue(codecURIEven, []byte("https://"+host+"/wiki/"+title)), true
}
