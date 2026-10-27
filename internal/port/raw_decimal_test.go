//ff:func feature=port type=helper control=iteration dimension=1
//ff:what TestRawDecimal — zero, PHP's exponent form with its extra 0, plain fractions, integers, mixed numbers
package port

import (
	"math/big"
	"testing"
)

func TestRawDecimal(t *testing.T) {
	r := func(s string) *big.Rat { v, _ := new(big.Rat).SetString(s); return v }
	for in, want := range map[string]string{
		"0": "0", "1e-6": "0.0000010", "2.5e-5": "0.000025", "1/3600": "0.00027777777777778",
		"1/60": "0.016666666666667", "100": "100", "12.5": "12.5", "1234567890123456": "1234567890123500",
	} {
		if got := rawDecimal(r(in)); got != want {
			t.Fatalf("%s: %q want %q", in, got, want)
		}
	}
}
