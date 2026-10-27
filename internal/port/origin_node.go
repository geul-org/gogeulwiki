//ff:func feature=port type=encoder control=sequence
//ff:what originNode — the LOCAL standing for a described source: the one declared before for the same description, else a new one bound to its URL or described by TRIPLEs
package port

import "github.com/geul-org/gogeul/pkg/core"

// originNode returns the LOCAL origin for describe (sorted by property,
// sortedSnaks) and the records declaring it. A description declared before
// in this document (st.declared, keyed by snakSet) reuses that LOCAL with no
// new records. Otherwise a URL-only description (urlOnly) binds a LOCAL to
// the URI by REFERENCE-NODE (localDefinition), and any other is a LOCAL node
// with one TRIPLE(P, LOCAL, value, normal) per snak (describedNode).
func originNode(a IDs, describe []refSnak, st *sourcing, rep *Report) (core.Record, []core.Record) {
	sorted := sortedSnaks(describe)
	key := snakSet(snakTexts(sorted))
	if local, ok := st.declared[key]; ok {
		return core.EncodeLocalRef(local), nil
	}
	var decl []core.Record
	var local uint64
	if urlOnly(sorted) {
		var def core.Record
		def, local = localDefinition(a, sorted[0].Value)
		decl = []core.Record{def}
	} else {
		decl, local = describedNode(a, sorted)
	}
	st.declared[key] = local
	rep.OriginLocals++
	rep.OriginEdges += len(decl)
	return core.EncodeLocalRef(local), decl
}
