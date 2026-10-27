//ff:func feature=port type=helper control=iteration dimension=1
//ff:what splitSnaks — a reference's snaks into those describing the source (originProperties) and those describing the citation, each in order
package port

// splitSnaks puts each snak of ss whose property is in originProperties in
// describe, every other one in cite, keeping the order of ss.
func splitSnaks(ss []refSnak) (describe, cite []refSnak) {
	for _, s := range ss {
		if originProperties[s.Property] {
			describe = append(describe, s)
		} else {
			cite = append(cite, s)
		}
	}
	return describe, cite
}
