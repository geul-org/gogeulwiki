//ff:func feature=port type=encoder control=sequence
//ff:what TestRevisionURI — the oldid permanent link of a revision
package port

import "testing"

func TestRevisionURI(t *testing.T) {
	if got := revisionURI(2557857485); got != "https://www.wikidata.org/w/index.php?oldid=2557857485" {
		t.Fatal(got)
	}
}
