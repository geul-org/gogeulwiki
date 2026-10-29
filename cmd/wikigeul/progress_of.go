//ff:func feature=cli type=helper control=sequence
//ff:what progressOf — the identity a pack-mode run checkpoints (commits, registry, dump, range or ID list, RECORDED-AT), nothing done yet
package main

// progressOf is the packProgress of run c with ID list hash ids, at the start.
func progressOf(c config, ids string) packProgress {
	return packProgress{Commit: c.Commit, LibCommit: c.LibCommit, LangDoc: c.LangDoc, Schema: c.Schema, Dump: c.Dump,
		From: c.From, To: c.To, IDs: ids, RecordedAt: c.RecordedAt}
}
