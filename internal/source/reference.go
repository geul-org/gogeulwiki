//ff:type feature=source type=model
//ff:what Reference — one Wikidata reference of a statement: statement, place in the statement's references, hash and its snaks in snaks-order
package source

// Reference is one reference (출처 묶음) of a statement of a loaded dump
// (sql/wikidata.sql statement_reference, reference_snak): the statement_id,
// its place among the statement's references (ordinal), the reference hash
// and its snaks in dump order (snaks-order), each in the same form as a
// qualifier (Qualifier: property, value type, datatype, value, snak type).
type Reference struct {
	StatementID int64
	Ordinal     int
	Hash        string
	Snaks       []Qualifier
}
