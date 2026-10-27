//ff:func feature=port type=helper control=sequence
//ff:what localOfParts — the LOCAL ID in part i of split Edge parts; false when there is no such part or it is not a LOCAL
package port

import "github.com/geul-org/gogeul/pkg/core"

// localOfParts is localOf of ps[i], false when ps has no part i.
func localOfParts(ps [][]core.Word, i int) (uint64, bool) {
	if i >= len(ps) {
		return 0, false
	}
	return localOf(ps[i])
}
