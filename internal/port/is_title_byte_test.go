//ff:func feature=port type=helper control=iteration dimension=1
//ff:what TestIsTitleByte — letters, digits and wfUrlencode's kept punctuation pass; % & ' + = ? # space, non-ASCII and NUL are encoded
package port

import "testing"

func TestIsTitleByte(t *testing.T) {
	for _, c := range []byte("aZ9-._;@$!*(),/~:") {
		if !isTitleByte(c) {
			t.Fatalf("%q: want kept", c)
		}
	}
	for _, c := range []byte{'%', '&', '\'', '+', '=', '?', '#', ' ', '"', 0xED, 0} {
		if isTitleByte(c) {
			t.Fatalf("%q: want encoded", c)
		}
	}
}
