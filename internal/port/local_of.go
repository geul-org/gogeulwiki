//ff:func feature=port type=helper control=sequence
//ff:what localOf — the LOCAL ID a record's words hold; false for any other record
package port

import "github.com/geul-org/gogeul/pkg/core"

// localOf decodes ws as one record and returns its LOCAL ID when it is a
// LOCAL Reference (Quantifier 000).
func localOf(ws []core.Word) (uint64, bool) {
	r, _, err := core.DecodeRecord(ws)
	if err != nil {
		return 0, false
	}
	id, err := core.DecodeLocalRef(r)
	return id, err == nil
}
