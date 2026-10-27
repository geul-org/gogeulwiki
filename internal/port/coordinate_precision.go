//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what coordinatePrecision — the uncertainty (binary32) and unit item a Wikidata coordinate precision text is written in without loss: degree, arcminute or arcsecond, the first whose precisionText gives the text back
package port

import (
	"math/big"

	"github.com/geul-org/gogeul/pkg/core"
)

// coordinatePrecision writes the precision text p (degrees, as the loaded
// value has it) as an uncertainty x in a unit (value/numeric/COORDINATE-EARTH
// 위키데이터 대응): for degree, arcminute and arcsecond in turn
// (precisionUnits), x is the binary32 nearest to p times the unit's count per
// degree, and the first unit whose precisionText(x) is p again is taken, so
// the conversion reads back to the same text. 10^k degrees become degrees
// (0.0001 → 0.0001°), 1/60° → 1 arcminute, 1/3600° → 1 arcsecond, 1/36000°
// → 0.1 arcsecond. It is false when p is not a number ≥ 0 or no unit gives
// it back.
func coordinatePrecision(p string) (float32, core.Record, bool) {
	v, ok := new(big.Rat).SetString(p)
	if !ok || v.Sign() < 0 {
		return 0, core.Record{}, false
	}
	for _, u := range precisionUnits {
		f, _ := new(big.Rat).Mul(v, big.NewRat(u.PerDeg, 1)).Float32()
		if precisionText(f, u.PerDeg) == p {
			unit, _ := entityRef(u.QID)
			return f, unit, true
		}
	}
	return 0, core.Record{}, false
}
