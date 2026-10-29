//ff:func feature=cli type=helper control=sequence
//ff:what runLine — the report's first line: the run's range, dump and commits as a # comment
package main

import "fmt"

// runLine is the comment that opens the report: the QID range, the loaded
// dump (schema, dump_id), the porter and library commits and the GEUL Language
// Registry version and content hash (LangDoc) of run c. The documents depend
// on the dump, the commits and that registry version only; the range says
// which documents this report covers.
func runLine(c config) string {
	return fmt.Sprintf("# wikigeul -from %d -to %d -schema %s -dump-id %d -commit %s -lib-commit %s -langreg %s", c.From, c.To, c.Schema, c.Dump, c.Commit, c.LibCommit, c.LangDoc)
}
