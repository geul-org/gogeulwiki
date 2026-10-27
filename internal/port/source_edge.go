//ff:func feature=port type=encoder control=sequence
//ff:what sourceEdge — a SOURCE Edge: Edge Id, Target (LOCAL Edge Id of the cited record, 0 for the document), Origin
package port

import "github.com/geul-org/gogeul/pkg/core"

// sourceEdge builds a SOURCE Edge (edge/SOURCE) tying the Edge target (0:
// the whole document, a HEAD SOURCE) to origin (a Reference or ABSENT) and
// returns it with its Edge Id.
func sourceEdge(a IDs, target uint64, origin core.Record) (core.Record, uint64) {
	return edgeWith(a, edgeSource, []core.Record{core.EncodeLocalRef(target), origin})
}
