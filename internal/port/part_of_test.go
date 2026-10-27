//ff:func feature=port type=helper control=sequence
//ff:what TestPartOf — an existing part, nil past the end
package port

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestPartOf(t *testing.T) {
	ps := [][]core.Word{{1}, {2}}
	if partOf(ps, 1)[0] != 2 || partOf(ps, 2) != nil {
		t.Fatal(ps)
	}
}
