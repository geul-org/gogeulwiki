//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what TestProvenanceOf — the dump's URL and download time, the converter and the library at their commits (not the loader); wikipg under its old names passes; another loader fails
package main

import (
	"testing"
	"time"

	"github.com/geul-org/wikigeul/internal/port"
	"github.com/geul-org/wikigeul/internal/source"
)

func TestProvenanceOf(t *testing.T) {
	d := source.Dump{URL: "u", SHA256: "s", DownloadedAt: time.Unix(5, 0), Loader: wikipgPath, LoaderCommit: "l"}
	p, err := provenanceOf(d, "abc", "def")
	if err != nil || p.DumpURI != "u" || p.DumpSHA256 != "s" || p.AcquiredAt != 5 || len(p.Tools) != 2 ||
		p.Tools[0] != (port.Tool{Doc: 1, Commit: "abc"}) || p.Tools[1] != (port.Tool{Doc: 2, Commit: "def"}) {
		t.Fatalf("got %+v %v", p, err)
	}
	for _, old := range []string{legacyWikiloadPath, legacyWikiportPath} {
		d.Loader = old
		if _, err := provenanceOf(d, "abc", "def"); err != nil {
			t.Fatalf("wikipg under its old name %s: %v", old, err)
		}
	}
	d.Loader = "github.com/geul-org/geul-entity/artifacts/wikidata"
	if _, err := provenanceOf(d, "abc", "def"); err == nil {
		t.Fatal("other loader: want error")
	}
}
