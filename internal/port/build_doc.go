//ff:func feature=port type=encoder control=sequence
//ff:what BuildDoc — the GEUL document of one Wikidata entity under issuer, with its report
package port

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"

	"github.com/geul-org/gogeulwiki/internal/source"
)

// BuildDoc writes the document of entity e under issuer (README): the
// HEAD provenance with its SOURCEs on the whole document (headRecords), then
// the BODY: the records declaring those SOURCEs' origins, LABEL per label, DESCRIPTION per
// description, ALIAS per alias, a property's datatype (datatypeEdges) and a
// TRIPLE followed by its QUALIFIERs and its references (SOURCE, README 출처)
// per convertible claim, then the sitelinks
// (sitelinkEdges). Edge Ids come from a:
// a statement equal to one of the previous version keeps its ID, new ones
// take the next (OVERVIEW 개정, delta.LocalIDs). The document ID is the
// DocNumber of e's ID (its letter and number). The report counts what was
// written and skipped, and the TRIPLEs whose references do not read back
// from the body as the same snak sets (RefMismatch).
func BuildDoc(e source.Entity, issuer []core.Word, p Provenance, a IDs) ([]byte, Report, error) {
	rep := Report{ID: e.ID, Skipped: map[string]int{}, Raw: map[string]int{}, Markers: map[string]int{}}
	prefix, n, okNum := entityNumber(e.ID)
	subject, okRef := entityRef(e.ID)
	if !okNum || !okRef {
		return nil, rep, fmt.Errorf("wikigeul: %q is not an item or property ID", e.ID)
	}
	num, err := DocNumber(prefix, n)
	if err != nil {
		return nil, rep, err
	}
	head, decl, err := headRecords(issuer, num, e, p, a)
	if err != nil {
		return nil, rep, err
	}
	labels := textEdges(edgeLabel, a, subject, e.Labels, &rep)
	descriptions := textEdges(edgeDescription, a, subject, e.Descriptions, &rep)
	aliases := textEdges(edgeAlias, a, subject, e.Aliases, &rep)
	body := append(append(append(decl, labels...), descriptions...), aliases...)
	body = append(body, datatypeEdges(a, subject, e.Datatype, &rep)...)
	st := newSourcing()
	body = append(body, tripleEdges(a, subject, e.Claims, st, &rep)...)
	body = append(body, sitelinkEdges(a, subject, e.Sitelinks, &rep)...)
	rep.Labels, rep.Descriptions, rep.Aliases = len(labels), len(descriptions), len(aliases)
	rep.RefMismatch = refMismatch(st.expect, referenceSnaks(body))
	return document(head, body), rep, nil
}
