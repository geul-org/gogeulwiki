//ff:func feature=port type=helper control=sequence
//ff:what TestQualifierClaim — property, type, value and snak type carry over unchanged
package port

import (
	"testing"

	"github.com/geul-org/gogeulwiki/internal/source"
)

func TestQualifierClaim(t *testing.T) {
	c := qualifierClaim(source.Qualifier{StatementID: 7, Property: "P512", Type: "wikibase-entityid", Value: "Q1765120", SnakType: "value"})
	if c.Property != "P512" || c.Type != "wikibase-entityid" || c.Value != "Q1765120" || c.SnakType != "value" || c.ID != 0 {
		t.Fatalf("got %+v", c)
	}
}
