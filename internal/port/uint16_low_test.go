//ff:func feature=port type=helper control=sequence
//ff:what uint16Low — test helper: the low WORD of a number, as the last payload WORD of a Reference to it
package port

import "github.com/geul-org/gogeul/pkg/core"

// uint16Low returns n's low 16 bits as a WORD.
func uint16Low(n uint64) core.Word {
	return core.Word(n)
}
