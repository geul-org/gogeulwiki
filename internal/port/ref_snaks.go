//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what refSnaks — the convertible snaks of a reference, in snaks-order; the others counted by reason with the prefix "reference "
package port

import "github.com/geul-org/wikigeul/internal/source"

// refSnaks converts each snak of r like a qualifier (qualifierClaim,
// claimObject with the registry langs; somevalue/novalue become slot markers). A snak whose property
// or value does not convert is left out and counted in rep.Skipped, dates
// kept as original text in rep.Raw and markers in rep.Markers, all prefixed
// "reference "; the properties used are added to rep.Properties.
func refSnaks(r source.Reference, langs Languages, rep *Report) []refSnak {
	var out []refSnak
	for _, q := range r.Snaks {
		c := qualifierClaim(q)
		prop, okProp := entityRef(q.Property)
		value, reason, raw := claimObject(c, langs)
		if !okProp {
			reason = "속성 ID: " + q.Property
		}
		if reason != "" {
			noteSkip(rep, "reference "+reason, q.Property, q.Value)
			continue
		}
		if raw != "" {
			noteRaw(rep, "reference "+raw, q.Property, q.Value)
		}
		rep.Properties = append(rep.Properties, propertiesOf(c)...)
		countMarker(rep, "reference ", value)
		out = append(out, refSnak{Property: q.Property, Prop: prop, Value: value})
	}
	return out
}
