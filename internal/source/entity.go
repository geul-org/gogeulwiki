//ff:type feature=source type=model
//ff:what Entity — one Wikidata item or property with its revision, labels, descriptions, aliases in every language, its claims and its sitelinks
package source

// Entity is one Wikidata item (Q) or property (P) of a loaded dump: a
// property's datatype (url, external-id …; empty for an item), the revision
// the dump holds (LastRevID, 0 when the dump gives none) and when it was
// saved (Modified, POSIX seconds, 0 when unknown),
// labels, descriptions and aliases in every language of the dump (mul
// included; aliases in dump order) and
// every claim in statement ID order (between properties the property ID as
// text, as wikipg numbers them; within a property dump order — Wikidata
// gives the order between properties no meaning) and every sitelink by site
// ID.
type Entity struct {
	ID           string
	Datatype     string
	LastRevID    int64
	Modified     int64
	Labels       []Text
	Descriptions []Text
	Aliases      []Text
	Claims       []Claim
	Sitelinks    []Sitelink
}
