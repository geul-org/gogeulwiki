//ff:type feature=port type=model
//ff:what sourcing — the references of one document being built: LOCAL origins declared so far by description, and the snak sets each TRIPLE should read back
package port

// sourcing is per-document state of sourceEdges. declared maps the
// description of a LOCAL origin (snakSet of its sorted describing snaks) to
// the LOCAL that stands for it, so equal sources are declared once per
// document. expect maps each TRIPLE Edge Id to the snak set of each of its
// references as converted (snakSet), for the round trip (referenceSnaks).
type sourcing struct {
	declared map[string]uint64
	expect   map[uint64][]string
}
