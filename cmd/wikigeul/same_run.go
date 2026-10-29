//ff:func feature=cli type=helper control=sequence
//ff:what sameRun — whether a checkpoint was made by the same run (commits, registry, dump, range or ID list, RECORDED-AT)
package main

// sameRun compares the identity fields of a resumed checkpoint p with the run
// now (want); progress fields (Kind, Num, Docs, Report) are not compared.
func sameRun(p, want packProgress) bool {
	want.Kind, want.Num, want.Docs, want.Report = p.Kind, p.Num, p.Docs, p.Report
	return p == want
}
