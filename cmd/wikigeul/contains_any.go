//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what containsAny — whether s contains any of the parts
package main

import "strings"

// containsAny reports whether s contains one of parts.
func containsAny(s string, parts []string) bool {
	for _, p := range parts {
		if strings.Contains(s, p) {
			return true
		}
	}
	return false
}
