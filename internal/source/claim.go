//ff:type feature=source type=model
//ff:what Claim — one Wikidata statement of a loaded dump: ID, property, value type, datatype, value, rank, snak type, its qualifiers and references
package source

// Claim is one statement of a loaded dump (sql/wikidata.sql statement): its
// statement_id, the property of the main snak, the datavalue type
// (wikibase-entityid, string, time, monolingualtext, globecoordinate,
// quantity …), the property's datatype (url, commonsMedia, external-id …;
// a string value carries several), the value (queriesFor: a bare entity ID, an unquoted string or
// the datavalue JSON; empty for novalue and somevalue), the rank (preferred,
// normal, deprecated), the main snak's type (value, somevalue, novalue) and
// its qualifiers and references in dump order.
type Claim struct {
	ID         int64
	Property   string
	Type       string
	Datatype   string
	Value      string
	Rank       string
	SnakType   string
	Qualifiers []Qualifier
	References []Reference
}
