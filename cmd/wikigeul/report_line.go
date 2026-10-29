//ff:func feature=cli type=helper control=sequence
//ff:what reportLine — one TSV line of a report: ID, status (ok, missing), label, description, alias, triple, qualifier, sitelink and badge counts, skipped and raw-kept reasons, slot markers written, with counts; reference counts
package main

import (
	"fmt"
	"strings"

	"github.com/geul-org/wikigeul/internal/port"
)

// reportColumns is the number of columns after ID and status (writeReport's
// header).
const reportColumns = 17

// reportHeader is the report's column line, after the run line.
const reportHeader = "id\tstatus\tlabels\tdescriptions\taliases\ttriples\tqualifiers\tsitelinks\tbadges\tskipped\traw\tmarkers\tsources\torigin_qid\torigin_absent\torigin_local\torigin_edges\tsource_qualifiers\tref_mismatch"

// reportLine formats r as ID, status, labels, descriptions, aliases, triples,
// qualifiers, sitelinks, badges and the skipped reasons, then the reasons of
// dates kept as original text, then the slot markers written for
// somevalue/novalue (reasonCounts), then the reference counts (SOURCE Edges,
// origins as item·ABSENT·LOCAL, the records declaring LOCAL origins,
// QUALIFIERs on SOURCEs, round-trip mismatches), tab-separated. The status
// is ok for a written document; an entity the dump does not have (r.Missing,
// no document) is "missing" with every other column empty, so it is not
// read as an empty entity.
func reportLine(r port.Report) string {
	if r.Missing {
		return r.ID + "\tmissing" + strings.Repeat("\t", reportColumns)
	}
	return fmt.Sprintf("%s\tok\t%d\t%d\t%d\t%d\t%d\t%d\t%d\t%s\t%s\t%s\t%d\t%d\t%d\t%d\t%d\t%d\t%d", r.ID, r.Labels, r.Descriptions, r.Aliases, r.Triples, r.Qualifiers, r.Sitelinks, r.Badges, reasonCounts(r.Skipped), reasonCounts(r.Raw), reasonCounts(r.Markers),
		r.Sources, r.OriginQIDs, r.OriginNone, r.OriginLocals, r.OriginEdges, r.SourceQuals, r.RefMismatch)
}
