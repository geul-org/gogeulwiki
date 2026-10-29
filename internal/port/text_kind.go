//ff:func feature=port type=helper control=selection
//ff:what textKind — the record name of a text Edge codec (LABEL, DESCRIPTION, ALIAS) for a Note
package port

// textKind names the text Edge codec for a Note's property: LABEL,
// DESCRIPTION or ALIAS, and "text" for any other codec.
func textKind(codec uint8) string {
	switch codec {
	case edgeLabel:
		return "LABEL"
	case edgeDescription:
		return "DESCRIPTION"
	case edgeAlias:
		return "ALIAS"
	default:
		return "text"
	}
}
