//ff:func feature=port type=helper control=selection
//ff:what pow10Rat — 10 to an integer power as an exact rational
package port

import "math/big"

// pow10Rat returns 10^k exactly, for negative k too.
func pow10Rat(k int) *big.Rat {
	p := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(max(k, -k))), nil)
	switch {
	case k < 0:
		return new(big.Rat).SetFrac(big.NewInt(1), p)
	default:
		return new(big.Rat).SetInt(p)
	}
}
