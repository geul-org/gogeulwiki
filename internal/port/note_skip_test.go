//ff:func feature=port type=helper control=sequence
//ff:what TestNoteSkip — the reason is counted in Skipped and a skipped Note keeps property and value
package port

import "testing"

func TestNoteSkip(t *testing.T) {
	rep := Report{Skipped: map[string]int{}}
	noteSkip(&rep, "개체: L", "P5137", "L7")
	noteSkip(&rep, "개체: L", "P5137", "L8")
	if rep.Skipped["개체: L"] != 2 || len(rep.Notes) != 2 || rep.Notes[1] != (Note{Kind: "skipped", Reason: "개체: L", Property: "P5137", Value: "L8"}) {
		t.Fatalf("%v %+v", rep.Skipped, rep.Notes)
	}
}
