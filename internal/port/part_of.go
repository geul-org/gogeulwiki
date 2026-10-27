//ff:func feature=port type=helper control=sequence
//ff:what partOf — part i of split Edge parts, or nothing when there is no such part
package port

import "github.com/geul-org/gogeul/pkg/core"

// partOf returns ps[i], nil when ps has no part i.
func partOf(ps [][]core.Word, i int) []core.Word {
	if i >= len(ps) {
		return nil
	}
	return ps[i]
}
