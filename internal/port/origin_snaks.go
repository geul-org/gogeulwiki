//ff:func feature=port type=helper control=selection
//ff:what originSnaks — the reference snaks an Origin stands for: P248=item for a WIKIDATA-QID, the LOCAL's description for a LOCAL, none for ABSENT or UNKNOWN-VALUE
package port

import "github.com/geul-org/gogeul/pkg/core"

// originSnaks inverts originOf for the Origin words o: a WIKIDATA-QID is the
// single snak P248 (stated in) = that item, a LOCAL is what desc (describedBy)
// says of it, and ABSENT (no description; UNKNOWN-VALUE alike) is no snak.
func originSnaks(o []core.Word, desc map[uint64][]string) []string {
	r, _, _ := core.DecodeRecord(o)
	local, isLocal := localOf(o)
	switch {
	case r.Kind == core.KindReference && r.Codec>>3 == core.RefWikidataQID:
		p248, _ := entityRef("P248")
		return []string{snakText(wordsOf(p248), o)}
	case isLocal:
		return desc[local]
	default:
		return nil
	}
}
