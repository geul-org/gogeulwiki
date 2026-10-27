//ff:type feature=port type=model
//ff:what refSnak — one converted reference snak: its Wikidata property ID, the property Reference and the value record (or slot marker)
package port

import "github.com/geul-org/gogeul/pkg/core"

// refSnak is a reference snak after conversion (refSnaks): the property ID as
// in the dump (P854), its WIKIDATA-PID Reference and the value record, a
// slot marker for somevalue/novalue.
type refSnak struct {
	Property string
	Prop     core.Record
	Value    core.Record
}
