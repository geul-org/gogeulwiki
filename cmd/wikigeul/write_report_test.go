//ff:func feature=cli type=helper control=sequence
//ff:what TestWriteReport — the run comment, a header and one line per report; the header has as many columns as a row
package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/geul-org/gogeulwiki/internal/port"
)

func TestWriteReport(t *testing.T) {
	p := filepath.Join(t.TempDir(), "r.tsv")
	if err := writeReport(p, "# run", []port.Report{{ID: "Q1"}, {ID: "Q2"}}); err != nil {
		t.Fatal(err)
	}
	b, _ := os.ReadFile(p)
	if lines := strings.Split(strings.TrimSpace(string(b)), "\n"); len(lines) != 4 || lines[0] != "# run" || !strings.HasPrefix(lines[1], "id\t") || !strings.HasPrefix(lines[3], "Q2\t") ||
		strings.Count(lines[1], "\t") != strings.Count(lines[2], "\t") {
		t.Fatalf("got %q", b)
	}
}
