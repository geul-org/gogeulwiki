//ff:func feature=source type=loader control=sequence
//ff:what TestLoadReferences — Q42's references grouped in order with snaks in snaks-order (bare entity ID, string, somevalue), another dump's hash left out; a closed handle and wrong columns error
package source

import "testing"

func TestLoadReferences(t *testing.T) {
	db, s := fixtureDB(t)
	q := queriesFor(s).References
	rs, err := loadReferences(db, q, "Q42")
	if err != nil || len(rs) != 2 || rs[0].Ordinal != 0 || rs[0].Hash != "h1" || len(rs[0].Snaks) != 2 || rs[0].Snaks[0].Property != "P248" || rs[0].Snaks[0].Value != "Q5375741" ||
		rs[0].Snaks[1].Value != "biography/Douglas-Adams" || rs[0].Snaks[1].StatementID != 4201 || rs[1].Hash != "h2" || len(rs[1].Snaks) != 2 || rs[1].Snaks[1].SnakType != "somevalue" || rs[1].Snaks[1].Value != "" {
		t.Fatalf("got %+v %v", rs, err)
	}
	if _, err := loadReferences(closedDB(t), q, "Q42"); err == nil {
		t.Fatal("closed: want error")
	}
	if _, err := loadReferences(db, "SELECT 1, 2 FROM "+s.Schema+".statement WHERE entity_id = $1", "Q42"); err == nil {
		t.Fatal("columns: want error")
	}
}
