//ff:func feature=port type=helper control=sequence
//ff:what TestSigDigits — rounding to n digits half to even, trailing zeros dropped, a carry into a new digit, exponents above and below 1
package port

import (
	"math/big"
	"testing"
)

func TestSigDigits(t *testing.T) {
	if d, e := sigDigits(big.NewRat(1, 60), 14); d != "16666666666667" || e != -2 {
		t.Fatalf("1/60: %s %d", d, e)
	}
	if d, e := sigDigits(big.NewRat(125, 1000), 2); d != "12" || e != -1 {
		t.Fatalf("0.125 half to even: %s %d", d, e)
	}
	if d, e := sigDigits(big.NewRat(135, 1000), 2); d != "14" || e != -1 {
		t.Fatalf("0.135 half to even up: %s %d", d, e)
	}
	if d, e := sigDigits(big.NewRat(9999, 1), 2); d != "1" || e != 4 {
		t.Fatalf("9999 carry: %s %d", d, e)
	}
	if d, e := sigDigits(big.NewRat(1500, 1), 3); d != "15" || e != 3 {
		t.Fatalf("1500: %s %d", d, e)
	}
}
