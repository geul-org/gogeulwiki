//ff:func feature=port type=helper control=sequence
//ff:what noteRaw — count a value kept as original text in rep.Raw by reason and keep its Note (property, original value)
package port

// noteRaw counts one value kept as original text for reason in rep.Raw and
// keeps its Note (Kind "raw") with the property and the original value.
func noteRaw(rep *Report, reason, property, value string) {
	rep.Raw[reason]++
	rep.Notes = append(rep.Notes, Note{Kind: "raw", Reason: reason, Property: property, Value: value})
}
