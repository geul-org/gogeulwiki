//ff:func feature=port type=helper control=selection
//ff:what rawDecimal — a non-negative number written as Wikibase's float text (14 significant digits) after jsonb: plain decimal, with the extra 0 of PHP's exponent form ("1.0e-6" → 0.0000010)
package port

import (
	"math/big"
	"strings"
)

// rawDecimal writes v ≥ 0 rounded to 14 significant digits (sigDigits) as
// PostgreSQL jsonb prints the number PHP wrote: PHP uses an exponent below
// 1e-4 and then shows at least one fraction digit ("1.0e-6"), which jsonb
// keeps as a trailing 0 of the plain decimal (0.0000010); from 1e-4 up it is
// a plain decimal without trailing zeros; zero is 0.
func rawDecimal(v *big.Rat) string {
	if v.Sign() == 0 {
		return "0"
	}
	d, e := sigDigits(v, 14)
	switch {
	case e < -4 && len(d) == 1:
		return "0." + strings.Repeat("0", -e-1) + d + "0"
	case e < 0:
		return "0." + strings.Repeat("0", -e-1) + d
	case len(d) <= e+1:
		return d + strings.Repeat("0", e+1-len(d))
	default:
		return d[:e+1] + "." + d[e+1:]
	}
}
