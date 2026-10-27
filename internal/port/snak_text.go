//ff:func feature=port type=helper control=sequence
//ff:what snakText — one snak as text for comparing: property words = value words, in hex
package port

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// snakText writes a snak (a property record and a value record, as words)
// as one comparable string. Two snaks are the same when their texts are.
func snakText(prop, value []core.Word) string {
	return fmt.Sprintf("%04x=%04x", prop, value)
}
