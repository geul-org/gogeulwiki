//ff:func feature=cli type=helper control=sequence
//ff:what TestContainsAny — true when a part occurs, false when none or no parts
package main

import "testing"

func TestContainsAny(t *testing.T) {
	if !containsAny("time: JSON 아님", []string{"x", "JSON"}) || containsAny("abc", []string{"d"}) || containsAny("abc", nil) {
		t.Fatal("containsAny")
	}
}
