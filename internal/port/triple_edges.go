//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what tripleEdges — a TRIPLE and then its QUALIFIERs and SOURCEs per convertible claim, Edge Ids from next; skipped claims counted by reason
package port

import (
	"github.com/geul-org/gogeul/pkg/core"

	"github.com/geul-org/wikigeul/internal/source"
)

// tripleEdges builds a TRIPLE on subject for each claim whose predicate,
// value and rank convert, each followed by its QUALIFIERs (qualifierEdges)
// and its references (sourceEdges with the document's state st),
// numbering Edge Ids from *next, languages by the registry langs. Every other claim is
// counted in rep.Skipped by reason, a date kept as original text in rep.Raw
// by its conversion reason (claimObject), a somevalue or novalue object written
// as its slot marker counted in rep.Markers, and every property used as a
// predicate or a value is added to rep.Properties.
func tripleEdges(a IDs, subject core.Record, claims []source.Claim, st *sourcing, langs Languages, rep *Report) []core.Record {
	var out []core.Record
	for _, c := range claims {
		pred, okPred := entityRef(c.Property)
		obj, reason, raw := claimObject(c, langs)
		rank, okRank := rankWord(c.Rank)
		switch {
		case !okPred:
			reason = "속성 ID: " + c.Property
		case reason == "" && !okRank:
			reason = "rank: " + c.Rank
		}
		if reason != "" {
			noteSkip(rep, reason, c.Property, c.Value)
			continue
		}
		if raw != "" {
			noteRaw(rep, raw, c.Property, c.Value)
		}
		rep.Properties = append(rep.Properties, propertiesOf(c)...)
		countMarker(rep, "", obj)
		r, id := tripleEdge(a, pred, subject, obj, rank)
		out = append(out, r)
		rep.Triples++
		out = append(out, qualifierEdges(a, id, c.Qualifiers, langs, rep)...)
		out = append(out, sourceEdges(a, id, c.References, st, langs, rep)...)
	}
	return out
}
