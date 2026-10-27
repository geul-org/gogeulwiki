//ff:func feature=port type=helper control=sequence
//ff:what TestSnakTexts — one text per snak in order
package port

import "testing"

func TestSnakTexts(t *testing.T) {
	p, _ := entityRef("P813")
	q, _ := entityRef("Q1")
	got := snakTexts([]refSnak{{Prop: p, Value: q}, {Prop: q, Value: p}})
	if len(got) != 2 || got[0] != snakText(wordsOf(p), wordsOf(q)) || got[1] != snakText(wordsOf(q), wordsOf(p)) {
		t.Fatalf("%v", got)
	}
}
