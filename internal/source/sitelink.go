//ff:type feature=source type=model
//ff:what Sitelink — one sitelink of an entity: the Wikimedia site ID, the page title and its badge items
package source

// Sitelink is one sitelink of a loaded entity (sql/wikidata.sql sitelink):
// the site's global ID (enwiki, commonswiki, zh_min_nanwikisource …), the
// page title as Wikidata stores it (spaces, not underscores) and the badge
// item IDs (Q17437796 featured article …) in dump order.
type Sitelink struct {
	Site   string
	Title  string
	Badges []string
}
