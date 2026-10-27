//ff:func feature=source type=loader control=selection
//ff:what exactValue — a column value, or the exact one from its *_json column when the loader had to keep it there (\u0000)
package source

import "encoding/json"

// exactValue returns v unless vj, the *_json column wikipg fills when text
// or jsonb cannot hold a value (an escape \u0000), is set. A JSON string in vj
// gives its decoded text (NUL included), the way value #>> '{}' unquotes a
// string; any other JSON (an object such as a monolingual text) is returned as
// written, the way a jsonb value is.
func exactValue(v, vj string) string {
	var s string
	switch {
	case vj == "":
		return v
	case json.Unmarshal([]byte(vj), &s) == nil:
		return s
	default:
		return vj
	}
}
