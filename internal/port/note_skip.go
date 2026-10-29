//ff:func feature=port type=helper control=sequence
//ff:what noteSkip — count a left-out value in rep.Skipped by reason and keep its Note (property, original value)
package port

// noteSkip counts one skipped occurrence of reason in rep.Skipped and keeps
// its Note (Kind "skipped") with the property and the original value.
func noteSkip(rep *Report, reason, property, value string) {
	rep.Skipped[reason]++
	rep.Notes = append(rep.Notes, Note{Kind: "skipped", Reason: reason, Property: property, Value: value})
}
