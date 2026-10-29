//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what sitelinkEdges — per sitelink a page LOCAL (REFERENCE-NODE → page URI) and TRIPLE(schema:about, page, item), then its badges; unknown sites counted
package port

import (
	"github.com/geul-org/gogeul/pkg/core"

	"github.com/geul-org/wikigeul/internal/source"
)

// sitelinkEdges writes the sitelinks of subject the way Wikidata's RDF does
// (README sitelink): each page is a LOCAL bound to its URL (sitelinkURI,
// localDefinition) and TRIPLE(about, page, subject, normal) says the page
// is about the entity, about being a LOCAL bound to http://schema.org/about
// once per document, before the first page. The page's badges follow
// (badgeEdges). A sitelink of a site with no known host is counted in
// rep.Skipped; rep.Sitelinks counts the pages written.
func sitelinkEdges(a IDs, subject core.Record, links []source.Sitelink, rep *Report) []core.Record {
	var out []core.Record
	var about, badge uint64
	for _, l := range links {
		page, ok := sitelinkURI(l)
		if !ok {
			noteSkip(rep, "sitelink 사이트: "+l.Site, "sitelink", l.Site+" "+l.Title)
			continue
		}
		if about == 0 {
			var def core.Record
			def, about = localDefinition(a, core.EncodeBytesValue(codecURIEven, []byte(schemaOrg+"about")))
			out = append(out, def)
		}
		def, local := localDefinition(a, page)
		triple, _ := tripleEdge(a, core.EncodeLocalRef(about), core.EncodeLocalRef(local), subject, 0)
		out = append(out, def, triple)
		rep.Sitelinks++
		out = append(out, badgeEdges(a, local, l.Badges, &badge, rep)...)
	}
	return out
}
