//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestParseFlags — defaults and overrides (registry root and version too); a missing or short commit, a bad range, an unknown flag (-langs is gone) a registry version over 32 bits, an unknown -out and zero workers or checkpoint fail
package main

import (
	"strings"
	"testing"
)

func TestParseFlags(t *testing.T) {
	a, b := strings.Repeat("a", 40), strings.Repeat("b", 40)
	c, err := parseFlags([]string{"-commit", a, "-lib-commit", b})
	if err != nil || c.Root != "." || c.From != 1 || c.To != 1000 || c.Commit != a || c.LibCommit != b || c.Schema != "wikidata" || c.Dump != 1 || c.LangRoot != "." || c.LangVer != 0 || c.Out != "pack" || c.Workers != defaultWorkers() || c.Every != 100000 || c.Samples != 2048 {
		t.Fatalf("defaults: %+v %v", c, err)
	}
	c, err = parseFlags([]string{"-commit", a, "-lib-commit", b, "-from", "5", "-to", "9", "-report", "r.tsv", "-root", "/x", "-schema", "wd", "-dump-id", "3", "-langreg-root", "/y", "-langreg-version", "2"})
	if err != nil || c.From != 5 || c.To != 9 || c.Report != "r.tsv" || c.Root != "/x" || c.Schema != "wd" || c.Dump != 3 || c.LangRoot != "/y" || c.LangVer != 2 {
		t.Fatalf("overrides: %+v %v", c, err)
	}
	c, err = parseFlags([]string{"-commit", a, "-lib-commit", b, "-strict", "-anomaly", "x.tsv", "-allow", "raw:time: 0년", "-allow", "warn:E7"})
	if err != nil || !c.Strict || c.Anomaly != "x.tsv" || len(c.Allow) != 2 || c.Allow[1] != "warn:E7" {
		t.Fatalf("strict: %+v %v", c, err)
	}
	if c, _ := parseFlags([]string{"-commit", a, "-lib-commit", b}); c.Strict || c.Anomaly != "wikigeul-anomaly.tsv" || c.Allow != nil {
		t.Fatalf("strict defaults: %+v", c)
	}
	for _, args := range [][]string{{}, {"-commit", a}, {"-commit", "abc", "-lib-commit", b}, {"-commit", a, "-lib-commit", b, "-from", "0"}, {"-commit", a, "-lib-commit", b, "-from", "9", "-to", "5"}, {"-nope"}, {"-commit", a, "-lib-commit", b, "-langs", "en"}, {"-commit", a, "-lib-commit", b, "-dump-id", "0"}, {"-commit", a, "-lib-commit", b, "-langreg-version", "4294967296"}, {"-commit", a, "-lib-commit", b, "-out", "zip"}, {"-commit", a, "-lib-commit", b, "-workers", "0"}, {"-commit", a, "-lib-commit", b, "-checkpoint", "0"}} {
		if _, err := parseFlags(args); err == nil {
			t.Fatalf("%v: want error", args)
		}
	}
}
