//ff:func feature=port type=helper control=sequence
//ff:what TestTextKind — each text Edge codec has its record name, any other codec is "text"
package port

import "testing"

func TestTextKind(t *testing.T) {
	got := []string{textKind(edgeLabel), textKind(edgeDescription), textKind(edgeAlias), textKind(edgeTriple)}
	if got[0] != "LABEL" || got[1] != "DESCRIPTION" || got[2] != "ALIAS" || got[3] != "text" {
		t.Fatalf("%v", got)
	}
}
