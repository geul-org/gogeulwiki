//ff:type feature=cli type=model
//ff:what config — the wikigeul run settings: tree root, QID range, DB DSN, loaded dump (and whether it keeps each reference occurrence's snak order), porter and library commits, report path
package main

// config is one wikigeul run: the archive root (geul-dev draft/geul-org/artifacts/geul) whose .well-known/geul receives
// the documents, the QID range [From, To] (every language is written: one
// document per entity holds them all), the PostgreSQL DSN (empty: lib/pq reads the PG* environment
// variables), the loaded dump to read (schema and dump_id, sql/wikidata.sql),
// the porter and GEUL library commits recorded in TRANSFORMED-BY and where
// the report goes. RefOrder is not a flag: run sets it from the loaded dump
// (source.Dump.RefOrder, reference_snak_order present).
type config struct {
	Root      string
	From      uint64
	To        uint64
	DSN       string
	Schema    string
	Dump      int16
	Commit    string
	LibCommit string
	Report    string
	RefOrder  bool
}
