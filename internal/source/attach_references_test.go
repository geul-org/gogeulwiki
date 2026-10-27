//ff:func feature=source type=loader control=sequence
//ff:what TestAttachReferences — references go to their claim in order; one of an unlisted claim is dropped
package source

import "testing"

func TestAttachReferences(t *testing.T) {
	cs := []Claim{{ID: 1}, {ID: 2}}
	attachReferences(cs, []Reference{{StatementID: 2, Hash: "a"}, {StatementID: 9, Hash: "x"}, {StatementID: 2, Hash: "b"}})
	if len(cs[0].References) != 0 || len(cs[1].References) != 2 || cs[1].References[0].Hash != "a" || cs[1].References[1].Hash != "b" {
		t.Fatalf("got %+v", cs)
	}
}
