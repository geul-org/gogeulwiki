//ff:func feature=cli type=helper control=sequence
//ff:what provenanceOf — the provenance every document of this run records: the loaded dump's URL, SHA-256 and time, and the converter and library commits
package main

import (
	"fmt"

	"github.com/geul-org/gogeulwiki/internal/port"
	"github.com/geul-org/gogeulwiki/internal/source"
)

// wikipgPath is the loader that wikipg (github.com/geul-org/wikipg) records in
// the dump row. The legacy paths are what it recorded under its earlier names
// (repository wikiport, command wikiload; the module was .../wikiport/go until
// 2026-10-27); dumps loaded then keep them.
const (
	wikipgPath         = "github.com/geul-org/wikipg/cmd/wikipg"
	legacyWikiloadPath = "github.com/geul-org/wikiport/cmd/wikiload"
	legacyWikiportPath = "github.com/geul-org/wikiport/go/cmd/wikiload"
)

// provenanceOf is the provenance of dump d (README 출처 기록): its URL,
// file SHA-256 and download time, then the tools as geul.org term definition
// documents (GEUL-ORG) — wikigeul (document 1, gogeulwiki) at commit and the
// GEUL library (document 2, gogeul) at libCommit. The loader is not a tool of
// the document: wikipg stores the dump losslessly (every line in
// entity_raw, values text or jsonb cannot hold in *_json), so the document
// follows from the dump and these two; its commit stays on the dump row. A
// loader other than wikipg is an error: its store is not known to be
// lossless.
func provenanceOf(d source.Dump, commit, libCommit string) (port.Provenance, error) {
	if d.Loader != wikipgPath && d.Loader != legacyWikiloadPath && d.Loader != legacyWikiportPath {
		return port.Provenance{}, fmt.Errorf("wikigeul: dump loaded by %q, not %s", d.Loader, wikipgPath)
	}
	return port.Provenance{
		DumpURI:    d.URL,
		DumpSHA256: d.SHA256,
		AcquiredAt: d.DownloadedAt.Unix(),
		Tools: []port.Tool{
			{Doc: 1, Commit: commit},
			{Doc: 2, Commit: libCommit},
		},
	}, nil
}
