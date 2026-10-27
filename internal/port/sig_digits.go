//ff:func feature=port type=helper control=iteration dimension=1
//ff:what sigDigits — the significant digits of a positive number rounded half to even to n digits, trailing zeros dropped, with its decimal exponent
package port

import (
	"math/big"
	"strings"
)

// sigDigits rounds v > 0 to n significant decimal digits (half to even, as
// a correctly rounded float formatter does) and returns those digits without
// trailing zeros and the exponent e of the first one (v ≈ 0.d × 10^(e+1)).
func sigDigits(v *big.Rat, n int) (string, int) {
	e := len(new(big.Int).Quo(v.Num(), v.Denom()).String()) - 1
	for t := new(big.Rat).Set(v); t.Cmp(big.NewRat(1, 1)) < 0; t.Mul(t, big.NewRat(10, 1)) {
		e--
	}
	scaled := new(big.Rat).Mul(v, pow10Rat(n-1-e))
	q, r := new(big.Int).QuoRem(scaled.Num(), scaled.Denom(), new(big.Int))
	if c := new(big.Int).Mul(r, big.NewInt(2)).Cmp(scaled.Denom()); c > 0 || (c == 0 && q.Bit(0) == 1) {
		q.Add(q, big.NewInt(1))
	}
	d := q.String()
	if len(d) > n {
		d, e = d[:n], e+1
	}
	return strings.TrimRight(d, "0"), e
}
