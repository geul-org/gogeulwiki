//ff:func feature=cli type=command control=iteration dimension=1
//ff:what parseFlags — the run config from command-line arguments; full porter and library commits are required
package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/geul-org/wikigeul/internal/port"
)

// parseFlags reads args into a config. -commit and -lib-commit (the full git
// object IDs of the porter and the GEUL library, written in every document's
// TRANSFORMED-BY) are required and must be well formed (port.CommitDigest).
// -schema and -dump-id name the loaded dump (default wikidata, 1).
func parseFlags(args []string) (config, error) {
	fs := flag.NewFlagSet("wikigeul", flag.ContinueOnError)
	var c config
	var dump int
	fs.StringVar(&c.Root, "root", ".", "archive root holding .well-known/geul (geul-dev draft/geul-org/artifacts/geul)")
	fs.Uint64Var(&c.From, "from", 1, "first QID number")
	fs.Uint64Var(&c.To, "to", 1000, "last QID number")
	fs.StringVar(&c.DSN, "dsn", "", "PostgreSQL DSN (empty: PG* environment variables)")
	fs.StringVar(&c.Schema, "schema", "wikidata", "schema of the loaded dump (wikipg -schema)")
	fs.IntVar(&dump, "dump-id", 1, "dump_id of the loaded dump (1..255)")
	fs.StringVar(&c.Commit, "commit", "", "porter git commit, full ID (required)")
	fs.StringVar(&c.LibCommit, "lib-commit", "", "GEUL library git commit, full ID (required)")
	fs.StringVar(&c.Report, "report", "", "report TSV path (empty: none)")
	fs.StringVar(&c.LangRoot, "langreg-root", "", "archive root to read the LANGUAGE-SHORT-CODE spec (GEUL Language Registry) from (empty: -root)")
	fs.StringVar(&c.Out, "out", "pack", "output: pack (the Issuer's pack set, zstd, streamed) or dir (directory tree, for tests)")
	fs.IntVar(&c.Workers, "workers", defaultWorkers(), "pack mode: conversion goroutines, each with its own database connection")
	fs.StringVar(&c.IDs, "ids", "", "pack mode: file of entity IDs (Q…/P…, one per line) to convert instead of every property and the Q range")
	fs.Int64Var(&c.Every, "checkpoint", 100000, "pack mode: documents between checkpoints")
	fs.BoolVar(&c.Resume, "resume", false, "pack mode: continue the interrupted run in -root from its last checkpoint")
	fs.Int64Var(&c.RecordedAt, "recorded-at", 0, "pack mode: RECORDED-AT of the version lists, Unix seconds (0: when the run starts)")
	fs.IntVar(&c.Samples, "samples", 2048, "pack mode: documents the zstd dictionary is trained on (the first Q documents of the run)")
	var langVer uint
	fs.UintVar(&langVer, "langreg-version", 0, "version of the LANGUAGE-SHORT-CODE spec document to use (0: the newest when the run starts)")
	if err := fs.Parse(args); err != nil {
		return config{}, err
	}
	for _, id := range []string{c.Commit, c.LibCommit} {
		if _, err := port.CommitDigest(id); err != nil {
			return config{}, fmt.Errorf("-commit and -lib-commit: %w", err)
		}
	}
	if dump < 1 || dump > 255 {
		return config{}, errors.New("wikigeul: need 1 <= -dump-id <= 255")
	}
	c.Dump = int16(dump)
	if langVer > 0xFFFFFFFF {
		return config{}, errors.New("wikigeul: -langreg-version is a 32-bit version ID")
	}
	c.LangVer = uint32(langVer)
	if c.Out != "pack" && c.Out != "dir" {
		return config{}, errors.New("wikigeul: -out is pack or dir")
	}
	if c.Workers < 1 || c.Every < 1 {
		return config{}, errors.New("wikigeul: -workers and -checkpoint must be at least 1")
	}
	if c.LangRoot == "" {
		c.LangRoot = c.Root
	}
	if c.From == 0 || c.From > c.To {
		return config{}, errors.New("wikigeul: need 1 <= -from <= -to")
	}
	return c, nil
}
