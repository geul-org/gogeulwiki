//ff:func feature=port type=encoder control=selection
//ff:what snakMarker — the slot marker of a special snak: somevalue → UNKNOWN-VALUE (0x0001), novalue → NO-VALUE (0x0002); false for a value snak
package port

import "github.com/geul-org/gogeul/pkg/core"

// snakMarker returns the GEUL slot marker (grammar reference/slot-marker) a
// Wikidata snak of type t stands for: somevalue (a value exists, unknown) is
// UNKNOWN-VALUE, novalue (no value) is NO-VALUE. A value snak (or an unknown
// type) has none and converts by its value.
func snakMarker(t string) (core.Record, bool) {
	var w core.Word
	switch t {
	case "somevalue":
		w = core.UnknownValue
	case "novalue":
		w = core.NoValue
	default:
		return core.Record{}, false
	}
	r, _, _ := core.DecodeRecord([]core.Word{w})
	return r, true
}
