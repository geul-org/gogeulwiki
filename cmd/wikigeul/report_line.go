//ff:func feature=cli type=helper control=sequence
//ff:what reportLine — one TSV line of a report: ID, label, description, alias, triple, qualifier, sitelink and badge counts, skipped and raw-kept reasons, slot markers written, with counts
package main

import (
	"fmt"

	"github.com/geul-org/gogeulwiki/internal/port"
)

// reportLine formats r as ID, labels, descriptions, aliases, triples, qualifiers, sitelinks, badges and the skipped
// reasons, then the reasons of dates kept as original text, then the slot
// markers written for somevalue/novalue (reasonCounts), tab-separated.
func reportLine(r port.Report) string {
	return fmt.Sprintf("%s\t%d\t%d\t%d\t%d\t%d\t%d\t%d\t%s\t%s\t%s", r.ID, r.Labels, r.Descriptions, r.Aliases, r.Triples, r.Qualifiers, r.Sitelinks, r.Badges, reasonCounts(r.Skipped), reasonCounts(r.Raw), reasonCounts(r.Markers))
}
