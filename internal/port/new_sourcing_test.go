//ff:func feature=port type=helper control=sequence
//ff:what TestNewSourcing — empty maps ready to write
package port

import "testing"

func TestNewSourcing(t *testing.T) {
	st := newSourcing()
	st.declared["a"], st.expect[1] = 2, []string{"x"}
	if len(st.declared) != 1 || len(st.expect) != 1 {
		t.Fatalf("%+v", st)
	}
}
