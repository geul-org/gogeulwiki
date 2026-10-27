//ff:func feature=port type=encoder control=sequence
//ff:what localDefinition — a REFERENCE-NODE defining a LOCAL by its Object (carried over by content), with that LOCAL
package port

import "github.com/geul-org/gogeul/pkg/core"

// localDefinition binds a LOCAL to object (edge/REFERENCE-NODE: Subject
// LOCAL → Object) with the Edge Id and LOCAL a gives for that object
// (IDs.Define), and returns the REFERENCE-NODE and the LOCAL.
func localDefinition(a IDs, object core.Record) (core.Record, uint64) {
	edge, local := a.Define(object)
	return compose(core.KindEdge, edgeReferenceNode, core.EncodeLocalRef(edge), core.EncodeLocalRef(local), object), local
}
