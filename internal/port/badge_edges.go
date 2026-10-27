//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what badgeEdges — TRIPLE(wikibase:badge, page, badge item) per badge of a sitelink, the badge predicate bound once; non-item badges counted
package port

import "github.com/geul-org/gogeul/pkg/core"

// badgeEdges writes a TRIPLE(badge, page, item, normal) for each badge item
// of the page LOCAL page, as Wikidata's RDF gives wikibase:badge on the page
// (README sitelink). *pred is the LOCAL bound to
// http://wikiba.se/ontology#badge; while it is 0 the binding is written
// first and *pred set, so a document binds it once. A badge that is not an
// item ID is counted in rep.Skipped; rep.Badges counts those written.
func badgeEdges(a IDs, page uint64, badges []string, pred *uint64, rep *Report) []core.Record {
	var out []core.Record
	for _, b := range badges {
		item, ok := entityRef(b)
		if !ok || b[0] != 'Q' {
			rep.Skipped["sitelink badge: "+b]++
			continue
		}
		if *pred == 0 {
			var def core.Record
			def, *pred = localDefinition(a, core.EncodeBytesValue(codecURIEven, []byte(wikibaseOntology+"badge")))
			out = append(out, def)
		}
		triple, _ := tripleEdge(a, core.EncodeLocalRef(*pred), core.EncodeLocalRef(page), item, 0)
		out = append(out, triple)
		rep.Badges++
	}
	return out
}
