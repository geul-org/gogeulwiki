//ff:type feature=port type=model
//ff:what Note — one value or text a document left out (skipped) or kept as original text (raw): reason, property and the original value
package port

// Note is one occurrence behind a Report count: Kind "skipped" (left out,
// counted in Report.Skipped) or "raw" (kept as original text, Report.Raw),
// the reason exactly as counted (with its "qualifier " or "reference "
// prefix), the property it sits on ("P31"; "LABEL", "DESCRIPTION" and
// "ALIAS" for texts, "sitelink", "badge", "datatype" for the rest) and the
// original value as the dump holds it. The strict run (wikigeul -strict)
// names each anomaly by these.
type Note struct {
	Kind     string
	Reason   string
	Property string
	Value    string
}
