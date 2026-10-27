//ff:type feature=source type=model
//ff:what Entity — one Wikidata item or property with its labels, descriptions, aliases in every language, its claims and its sitelinks
package source

// Entity is one Wikidata item (Q) or property (P) of a loaded dump: a
// property's datatype (url, external-id …; empty for an item),
// labels, descriptions and aliases in every language of the dump (mul
// included; aliases in dump order) and
// every claim in statement ID order and every sitelink by site ID.
type Entity struct {
	ID           string
	Datatype     string
	Labels       []Text
	Descriptions []Text
	Aliases      []Text
	Claims       []Claim
	Sitelinks    []Sitelink
}
