//ff:func feature=port type=encoder control=selection
//ff:what originOf — the Origin of a SOURCE from a reference's describing snaks: the P248 item alone, ABSENT when nothing describes it, else a LOCAL node (with the records declaring it the first time)
package port

import "github.com/geul-org/gogeul/pkg/core"

// originOf decides the Origin of a reference (README 출처, 원천 결정 규칙)
// from describe, its snaks of originProperties in dump order:
//   - nothing describes the source (only citing snaks such as P143, P813):
//     ABSENT, the origin is not written — the converter does not claim a
//     source exists that is not known (UNKNOWN-VALUE), it only has none to
//     write (reference/slot-marker ABSENT);
//   - exactly one snak, P248 (stated in) with an item: that WIKIDATA-QID;
//   - otherwise a LOCAL node the describing snaks declare (originNode),
//     P248 among them as TRIPLE P248.
//
// It returns the Origin record and the records to write before the SOURCE
// (empty when the Origin needs no declaration or was declared before).
func originOf(a IDs, describe []refSnak, st *sourcing, rep *Report) (core.Record, []core.Record) {
	switch {
	case len(describe) == 0:
		rep.OriginNone++
		return core.Record{Kind: core.KindReference, Codec: 0}, nil
	case len(describe) == 1 && describe[0].Property == "P248" && describe[0].Value.Kind == core.KindReference && describe[0].Value.Codec>>3 == core.RefWikidataQID:
		rep.OriginQIDs++
		return describe[0].Value, nil
	default:
		return originNode(a, describe, st, rep)
	}
}
