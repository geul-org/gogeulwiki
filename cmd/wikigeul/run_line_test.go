//ff:func feature=cli type=helper control=sequence
//ff:what TestRunLine — the comment names range, schema, dump and both commits
package main

import "testing"

func TestRunLine(t *testing.T) {
	got := runLine(config{From: 1, To: 10000, Schema: "wikidata", Dump: 1, Commit: "aa", LibCommit: "bb", LangDoc: "0x0000/0x0000000000000848@2 0x26fb"})
	if got != "# wikigeul -from 1 -to 10000 -schema wikidata -dump-id 1 -commit aa -lib-commit bb -langreg 0x0000/0x0000000000000848@2 0x26fb" {
		t.Fatalf("got %q", got)
	}
}
