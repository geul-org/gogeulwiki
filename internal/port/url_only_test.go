//ff:func feature=port type=helper control=iteration dimension=1
//ff:what TestURLOnly — one P854 URI is URL-only; a text URL, another property, two snaks are not
package port

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestURLOnly(t *testing.T) {
	uri := core.EncodeBytesValue(codecURIEven, []byte("http://x"))
	odd := core.EncodeBytesValue(codecURIEven, []byte("http://xy"))
	txt := core.EncodeBytesValue(core.CodecUTF8TextEven, []byte("x"))
	if !urlOnly([]refSnak{{Property: "P854", Value: uri}}) || !urlOnly([]refSnak{{Property: "P854", Value: odd}}) {
		t.Fatal("URI: want true")
	}
	for i, ss := range [][]refSnak{
		{{Property: "P854", Value: txt}},
		{{Property: "P1476", Value: uri}},
		{{Property: "P854", Value: uri}, {Property: "P854", Value: uri}},
		nil,
	} {
		if urlOnly(ss) {
			t.Fatalf("case %d: want false", i)
		}
	}
}
