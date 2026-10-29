//ff:func feature=cli type=command control=iteration dimension=1
//ff:what checkFacts — the Facts (Error, Warn) of layer-1 schema.Check over an entity's version 1 and version list
package main

import "github.com/geul-org/gogeul/pkg/schema"

// checkFacts runs schema.Check (layer 1, what `geul check` does to one
// document) on version and list and keeps every Fact (Error or Warn),
// naming which record each is on. A strict run (anomaliesOf) stops at any
// but the known Warn E6.
func checkFacts(version, list []byte) []checkFact {
	var out []checkFact
	for i, b := range [][]byte{version, list} {
		for _, f := range schema.Check(b, schema.Options{}).Facts {
			out = append(out, checkFact{Doc: []string{"version", "list"}[i], Fact: f})
		}
	}
	return out
}
