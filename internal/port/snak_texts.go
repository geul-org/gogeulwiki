//ff:func feature=port type=helper control=iteration dimension=1
//ff:what snakTexts — the texts of converted reference snaks, in order
package port

// snakTexts returns snakText of each snak of ss.
func snakTexts(ss []refSnak) []string {
	var out []string
	for _, s := range ss {
		out = append(out, snakText(wordsOf(s.Prop), wordsOf(s.Value)))
	}
	return out
}
