//ff:func feature=source type=loader control=sequence
//ff:what TestExactValue — no JSON keeps the value; a JSON string decodes with its NUL; other JSON is returned as written
package source

import "testing"

func TestExactValue(t *testing.T) {
	if got := exactValue("ab", ""); got != "ab" {
		t.Fatalf("plain: %q", got)
	}
	if got := exactValue("ab", `"a\u0000b"`); got != "a\x00b" {
		t.Fatalf("string: %q", got)
	}
	if got := exactValue("", `{"text":"a\u0000","language":"en"}`); got != `{"text":"a\u0000","language":"en"}` {
		t.Fatalf("object: %q", got)
	}
}
