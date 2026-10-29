//ff:func feature=port type=helper control=sequence
//ff:what TestNoteRaw — the reason is counted in Raw and a raw Note keeps property and value
package port

import "testing"

func TestNoteRaw(t *testing.T) {
	rep := Report{Raw: map[string]int{}}
	noteRaw(&rep, "qualifier time: 시간대·오차 범위", "P585", `{"time":"+2001"}`)
	if rep.Raw["qualifier time: 시간대·오차 범위"] != 1 || len(rep.Notes) != 1 || rep.Notes[0] != (Note{Kind: "raw", Reason: "qualifier time: 시간대·오차 범위", Property: "P585", Value: `{"time":"+2001"}`}) {
		t.Fatalf("%v %+v", rep.Raw, rep.Notes)
	}
}
