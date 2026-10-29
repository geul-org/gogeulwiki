//ff:func feature=cli type=helper control=sequence
//ff:what TestReportLine — ID, status ok, tab-separated counts (labels, descriptions, aliases, triples, qualifiers, sitelinks, badges), sorted skip reasons, raw-kept date reasons, slot markers, then the reference counts; a missing entity is status missing with empty columns
package main

import (
	"strings"
	"testing"

	"github.com/geul-org/wikigeul/internal/port"
)

func TestReportLine(t *testing.T) {
	got := reportLine(port.Report{ID: "Q1", Labels: 2, Descriptions: 6, Aliases: 3, Triples: 4, Qualifiers: 5, Sitelinks: 7, Badges: 1, Skipped: map[string]int{"b": 1, "a": 2}, Raw: map[string]int{"time: x": 1}, Markers: map[string]int{"qualifier UNKNOWN-VALUE": 2, "NO-VALUE": 1},
		Sources: 8, OriginQIDs: 1, OriginNone: 2, OriginLocals: 3, OriginEdges: 4, SourceQuals: 9, RefMismatch: 0})
	if got != "Q1\tok\t2\t6\t3\t4\t5\t7\t1\ta=2; b=1\ttime: x=1\tNO-VALUE=1; qualifier UNKNOWN-VALUE=2\t8\t1\t2\t3\t4\t9\t0" {
		t.Fatalf("got %q", got)
	}
	if got := reportLine(port.Report{ID: "Q6", Missing: true}); got != "Q6\tmissing"+strings.Repeat("\t", 17) {
		t.Fatalf("missing: %q", got)
	}
}
