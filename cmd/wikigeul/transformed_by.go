//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what transformedBy — a document's TRANSFORMED-BY HEAD records as one hex string ("" when it has none or does not split)
package main

import (
	"github.com/geul-org/gogeul/pkg/core"
)

// transformedBy encodes the HEAD Meta TRANSFORMED-BY (0x33) records of b in
// order, so documents of one run can be compared: they must name the same
// porter and library commits.
func transformedBy(b []byte) string {
	ws, _ := core.ReadWords(b)
	parts, err := core.SplitDocument(ws)
	if err != nil {
		return ""
	}
	var out []core.Word
	for _, r := range parts.Head {
		if r.Kind == core.KindMeta && r.Codec == 0x33 {
			w, _ := core.EncodeRecord(r)
			out = append(out, w...)
		}
	}
	return core.FormatID(out)
}
