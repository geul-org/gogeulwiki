//ff:type feature=cli type=model
//ff:what packProgress — what a pack-mode run saves with each stream checkpoint: the run's identity and how far it got
package main

// packProgress is saved (JSON) with every stream checkpoint
// (pack.Stream.Checkpoint meta). The identity fields — porter and library
// commits (TRANSFORMED-BY), registry version, dump, range or ID list, the
// RECORDED-AT of the version lists, the sampled dictionary — must match on
// resume, or the resumed documents would differ from the ones already
// written. Kind/Num is the last entity put, Docs how many documents, Report
// the report file's size at the checkpoint.
type packProgress struct {
	Commit     string
	LibCommit  string
	LangDoc    string
	Schema     string
	Dump       int16
	From       uint64
	To         uint64
	IDs        string
	RecordedAt int64
	Kind       byte
	Num        uint64
	Docs       int64
	Report     int64
}
