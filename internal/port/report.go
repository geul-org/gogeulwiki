//ff:type feature=port type=model
//ff:what Report — what one ported document holds and what it left out, with the properties it used
package port

// Report tells what one ported document holds and what it left out: counts
// of LABEL, DESCRIPTION, ALIAS, TRIPLE and QUALIFIER records, sitelink pages
// and their badges, skipped claims and
// qualifiers by reason (unsupported value type, Lexeme value …), dates
// kept as original text by conversion reason (Raw), somevalue/novalue snaks
// written as slot markers by marker name (Markers; qualifier ones prefixed
// "qualifier ") and the properties the
// document uses as predicates, qualifier properties or values, for porting
// those next.
type Report struct {
	ID           string
	Labels       int
	Descriptions int
	Aliases      int
	Triples      int
	Qualifiers   int
	Sitelinks    int
	Badges       int
	Skipped      map[string]int
	Raw          map[string]int
	Markers      map[string]int
	Properties   []string
}
