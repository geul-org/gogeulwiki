//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what writeReport — the run report as TSV: a # line naming the run, a header, one line per entity
package main

import (
	"os"
	"strings"

	"github.com/geul-org/wikigeul/internal/port"
)

// writeReport writes reps to path as TSV (reportLine) after the comment line
// head (runLine) and a header line.
func writeReport(path, head string, reps []port.Report) error {
	lines := []string{head, "id\tstatus\tlabels\tdescriptions\taliases\ttriples\tqualifiers\tsitelinks\tbadges\tskipped\traw\tmarkers\tsources\torigin_qid\torigin_absent\torigin_local\torigin_edges\tsource_qualifiers\tref_mismatch"}
	for _, r := range reps {
		lines = append(lines, reportLine(r))
	}
	return os.WriteFile(path, []byte(strings.Join(lines, "\n")+"\n"), 0o644)
}
