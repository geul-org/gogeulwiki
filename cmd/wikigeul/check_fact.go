//ff:type feature=cli type=model
//ff:what checkFact — a schema Fact (gogeul schema.Check) on one of the records an entity puts: its version 1 or its version list
package main

import "github.com/geul-org/gogeul/pkg/schema"

// checkFact is a schema.Check finding on Doc ("version" or "list").
type checkFact struct {
	Doc string
	schema.Fact
}
