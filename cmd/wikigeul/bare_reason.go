//ff:func feature=cli type=helper control=sequence
//ff:what bareReason — a report reason without its "qualifier " or "reference " prefix
package main

import "strings"

// bareReason drops the "qualifier " or "reference " prefix the report puts
// on reasons of qualifier and reference snaks, so one rule covers the main
// value, its qualifiers and its references.
func bareReason(reason string) string {
	reason = strings.TrimPrefix(reason, "qualifier ")
	return strings.TrimPrefix(reason, "reference ")
}
