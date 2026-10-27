//ff:func feature=port type=helper control=sequence
//ff:what countMarker — count a written slot marker in rep.Markers under its name, with a prefix ("" or "qualifier ")
package port

import "github.com/geul-org/gogeul/pkg/core"

// countMarker adds one to rep.Markers[prefix + marker name] when r is a slot
// marker (core.IsSlotMarker), creating the map on first use.
func countMarker(rep *Report, prefix string, r core.Record) {
	if !core.IsSlotMarker(r) {
		return
	}
	if rep.Markers == nil {
		rep.Markers = map[string]int{}
	}
	rep.Markers[prefix+core.SlotMarkerName(r.Codec)]++
}
