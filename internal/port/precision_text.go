//ff:func feature=port type=helper control=sequence
//ff:what precisionText — the Wikidata precision text an uncertainty x (binary32) in a unit stands for: x's shortest decimal turned into degrees, written as Wikibase writes a float (14 significant digits) and PostgreSQL jsonb reads it back
package port

import (
	"math/big"
	"strconv"
)

// precisionText inverts coordinatePrecision for x, an uncertainty in a unit
// perDeg of which make a degree: the shortest decimal that reads back as x
// (binary32), divided by perDeg exactly, is the precision in degrees. It is
// written as Wikibase serialized floats into the dump (PHP, 14 significant
// digits, exponent form below 1e-4: 0.016666666666667, 1.0e-6,
// 2.7777777777778e-5) and as the loader's jsonb then prints that number
// (plain decimal keeping the written digits: 0.0000010,
// 0.000027777777777778), which is the text the converter reads (rawDecimal).
func precisionText(x float32, perDeg int64) string {
	v, _ := new(big.Rat).SetString(strconv.FormatFloat(float64(x), 'g', -1, 32))
	return rawDecimal(new(big.Rat).Quo(v, big.NewRat(perDeg, 1)))
}
