//ff:func feature=port type=helper control=sequence
//ff:what wordsOf — a whole record's words (TAG and Length included)
package port

import "github.com/geul-org/gogeul/pkg/core"

// wordsOf encodes r whole. A record that does not encode gives no words.
func wordsOf(r core.Record) []core.Word {
	ws, _ := core.EncodeRecord(r)
	return ws
}
