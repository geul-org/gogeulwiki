//ff:type feature=cli type=model
//ff:what packJob — one entity to convert in pack mode, numbered in output order
package main

// packJob is entity Kind+Num (Kind 'P' or 'Q') at position Seq of the output
// order (document key order: every P before every Q, then by number).
type packJob struct {
	Seq  int64
	Kind byte
	Num  uint64
}
