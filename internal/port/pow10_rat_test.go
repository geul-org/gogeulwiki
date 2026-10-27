//ff:func feature=port type=helper control=sequence
//ff:what TestPow10Rat — positive, zero and negative powers of 10
package port

import (
	"math/big"
	"testing"
)

func TestPow10Rat(t *testing.T) {
	if pow10Rat(3).Cmp(big.NewRat(1000, 1)) != 0 || pow10Rat(0).Cmp(big.NewRat(1, 1)) != 0 || pow10Rat(-2).Cmp(big.NewRat(1, 100)) != 0 {
		t.Fatal("powers")
	}
}
