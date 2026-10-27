//ff:func feature=port type=helper control=sequence
//ff:what newSourcing — empty per-document reference state
package port

// newSourcing returns a sourcing with no origin declared and nothing expected.
func newSourcing() *sourcing {
	return &sourcing{declared: map[string]uint64{}, expect: map[uint64][]string{}}
}
