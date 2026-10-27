//ff:type feature=port type=model
//ff:what IDs — where a ported document gets its LOCAL IDs: the GEUL library's carried-over assignment (delta.LocalIDs)
package port

import "github.com/geul-org/gogeul/pkg/core"

// IDs hands out LOCAL IDs that carry over from the previous version of the
// document (OVERVIEW 개정: same statement same ID, new ones from the auto
// increment; a cited source node described by TRIPLEs keeps its LOCAL by
// what they say, Node). *delta.LocalIDs of the GEUL library implements it.
type IDs interface {
	Edge(codec uint8, rest []core.Word) uint64
	Define(object core.Record) (edge, local uint64)
	Node(about [][]core.Word) uint64
}
