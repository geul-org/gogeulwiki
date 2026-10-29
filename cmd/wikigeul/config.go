//ff:type feature=cli type=model
//ff:what config — the wikigeul run settings: tree root, QID range, DB DSN, loaded dump (and whether it keeps each reference occurrence's snak order), porter and library commits, report path
package main

import "github.com/geul-org/wikigeul/internal/port"

// config is one wikigeul run: the archive root (geul-dev draft/geul-org/artifacts/geul) whose .well-known/geul receives
// the documents, the QID range [From, To] (every language is written: one
// document per entity holds them all), the PostgreSQL DSN (empty: lib/pq reads the PG* environment
// variables), the loaded dump to read (schema and dump_id, sql/wikidata.sql),
// the porter and GEUL library commits recorded in TRANSFORMED-BY and where
// the report goes. RefOrder is not a flag: run sets it from the loaded dump
// (source.Dump.RefOrder, reference_snak_order present). LangRoot and LangVer
// say where and which version of the LANGUAGE-SHORT-CODE codec spec document
// to read the GEUL Language Registry from (LangRoot "" is Root, LangVer 0 the
// newest at the start of the run); run reads it once into Langs and LangDoc
// (languageRegistry), so every document of a run uses the same registry.
// Out is "pack" (default: the Issuer's pack set under Root written as one
// stream, runPack) or "dir" (the directory tree, for tests). Pack mode
// converts with Workers goroutines, the entities of IDs (a file of Q/P IDs)
// or else every property and Q<From>…Q<To>, checkpoints every Every
// documents, Resume continues an interrupted run, RecordedAt is the
// RECORDED-AT of every version list (0: the run's start) and Samples how many
// documents train the zstd dictionary. Strict (pack mode) stops the run at
// the first anomaly (anomaliesOf: a skip or raw-kept value outside the
// normal list, a failed conversion or panic, a missing entity, a reference
// mismatch, a schema Fact other than the known Warn E6), after a checkpoint,
// writing the anomaly report Anomaly; Allow adds "kind:reason" entries
// ("raw:time: 0년", "warn:E7") to the normal list for a resumed run.
type config struct {
	Root       string
	From       uint64
	To         uint64
	DSN        string
	Schema     string
	Dump       int16
	Commit     string
	LibCommit  string
	Report     string
	RefOrder   bool
	LangRoot   string
	LangVer    uint32
	Langs      port.Languages
	LangDoc    string
	Out        string
	Workers    int
	IDs        string
	Every      int64
	Resume     bool
	RecordedAt int64
	Samples    int
	Strict     bool
	Anomaly    string
	Allow      []string
}
