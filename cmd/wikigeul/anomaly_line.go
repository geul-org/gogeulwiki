//ff:func feature=cli type=helper control=sequence
//ff:what anomalyLine — one anomaly as a TSV line; value and reason Go-quoted so tabs and newlines stay on the line, a long value cut at 4096 bytes
package main

import (
	"strconv"
	"strings"
)

// anomalyLine formats a as id, kind, property, value, reason, class with
// tabs. The value (the original datavalue) and the reason (an error may
// carry a stack) are Go-quoted; a value over 4096 bytes is cut and marked
// "…(n bytes)".
func anomalyLine(a anomaly) string {
	v := a.Value
	if len(v) > 4096 {
		v = v[:4096] + "…(" + strconv.Itoa(len(a.Value)) + " bytes)"
	}
	return strings.Join([]string{a.ID, a.Kind, a.Property, strconv.Quote(v), strconv.Quote(a.Reason), a.Class}, "\t")
}
