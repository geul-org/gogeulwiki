//ff:type feature=source type=model
//ff:what entityQueries — the SQL LoadEntity runs: existence, labels, descriptions, aliases, claims, qualifiers and sitelinks of one entity
package source

// entityQueries is the SQL loadEntityWith runs for one entity ($1 = id;
// every language). LoadEntity uses queriesFor.
type entityQueries struct {
	Exists       string
	Labels       string
	Descriptions string
	Aliases      string
	Claims       string
	Qualifiers   string
	Sitelinks    string
}
