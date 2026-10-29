# wikigeul

`wikigeul` converts Wikidata entities into [GEUL](https://geul.org/) documents. It reads a Wikidata dump that [wikipg](https://github.com/geul-org/wikipg) has loaded into PostgreSQL and publishes one `.geul` document per item or property — labels, descriptions, aliases, claims with qualifiers and references, and sitelinks in every language — under `.well-known/geul/`.

All documents are published under one Issuer, `wikidata.geul.org` (`0x00010000`). The top byte of a document ID is the entity kind letter (`Q` = `0x51`, `P` = `0x50`) and the low 56 bits are its number, so Q42 is document `0x510000000000002a`.

## Install

```sh
go install github.com/geul-org/wikigeul/cmd/wikigeul@latest
```

## Requirements

- Go 1.25 or later.
- PostgreSQL holding a Wikidata dump loaded by [wikipg](https://github.com/geul-org/wikipg). wikigeul refuses a dump loaded by any other tool, because only wikipg's store is known to be lossless.
- The full git commit IDs of this tool and of the [gogeul](https://github.com/geul-org/gogeul) library it was built with; they are recorded in every document.
- An archive holding the LANGUAGE-SHORT-CODE codec spec document `0x0000/0x0000000000000848` (the GEUL Language Registry), by default `-root` itself.

## Usage

```sh
wikigeul -root ./archive -from 1 -to 1000 \
  -commit "$(git rev-parse HEAD)" \
  -lib-commit <gogeul commit> \
  -report report.tsv
```

By default (`-out pack`) the documents go straight into the Issuer's pack set `<root>/.well-known/geul/0x00010000/` (zstd, with a dictionary trained on the run's first Q documents): every property of the dump, then the dump's items Q<from>…Q<to> (or the entities of `-ids`), converted by `-workers` goroutines with their own database connections and written in document key order as one stream (gogeul `pack.Stream`). Each document is version 1 with its version list and is also served by content hash under `h/` without a copy. No `documents.geul` is written in pack mode: for all of Wikidata it would be a single 4 GB document, and a document's name follows from its ID. The run checkpoints every `-checkpoint` documents; `-resume` continues an interrupted run, which must be the same run (commits, registry version, dump, range or ID list, `-recorded-at`). The set must not exist yet: pack mode writes a whole new set.

With `-out dir` (for tests) it converts Q1–Q1000, then every property those documents use, into the directory tree `<root>/.well-known/geul/` and writes the Issuer's name index `documents.geul` (`q42`, `p31`, …). The index keeps documents already published there, so a range can be converted in several runs. Both modes give the same document bytes for the same input and commits.

The database connection comes from `-dsn` or, when it is empty, from the standard libpq environment variables (`PGHOST`, `PGPORT`, `PGUSER`, `PGDATABASE`, …) and the password file `~/.pgpass`.

| Flag | Meaning | Default |
|---|---|---|
| `-root` | Archive root; documents go under `<root>/.well-known/geul/` | `.` |
| `-from`, `-to` | QID range | `1`, `1000` |
| `-commit` | Full commit ID of wikigeul, recorded in TRANSFORMED-BY (required) | |
| `-lib-commit` | Full commit ID of the gogeul library, recorded in TRANSFORMED-BY (required) | |
| `-dsn` | PostgreSQL DSN | empty: `PG*` environment variables |
| `-schema` | Schema of the loaded dump (wikipg `-schema`) | `wikidata` |
| `-dump-id` | Which loaded dump to read (wikipg `-dump-id`, 1–255) | `1` |
| `-report` | Write a per-entity TSV report to this path | none |
| `-out` | `pack` (the Issuer's pack set, streamed) or `dir` (directory tree, for tests) | `pack` |
| `-workers` | Pack mode: conversion goroutines, each with a database connection | 70% of the CPUs halved (a database backend per worker) |
| `-ids` | Pack mode: file of entity IDs (`Q…`/`P…`, one per line) instead of every property and the Q range | none |
| `-checkpoint` | Pack mode: documents between checkpoints | `100000` |
| `-resume` | Pack mode: continue the interrupted run from its last checkpoint | off |
| `-recorded-at` | Pack mode: RECORDED-AT of every version list (Unix seconds) | the run's start |
| `-samples` | Pack mode: Q documents the zstd dictionary is trained on | `2048` |
| `-langreg-root` | Archive to read the GEUL Language Registry (spec document `0x0000/0x0000000000000848`) from | `-root` |
| `-langreg-version` | Version of that spec document to use | `0`: the newest when the run starts |

Language codes: each Wikimedia language code is mapped to a BCP 47 tag; when the GEUL Language Registry numbers the tag (compared without case) the text's language is written as a 1-WORD LANGUAGE-SHORT-CODE, otherwise as a LANGUAGE-BCP47-CODE in canonical case. The registry is not built into the code: it is read once per run from the spec document's ENUMs through the gogeul resolve path, and the version used (`0x0000/0x0000000000000848@n` and its content hash) is printed and recorded in the report's first line.

The report starts with a comment line recording the run (`# wikigeul -from … -to … -schema … -dump-id … -commit … -lib-commit … -langreg 0x0000/0x0000000000000848@n 0x<hash>`), then a header and one line per ID in the range. `status` is `ok` (document written) or `missing` (no such entity in the dump). The other columns count what was written and what was skipped: `labels`, `descriptions`, `aliases`, `triples`, `qualifiers`, `sitelinks`, `badges`, `skipped`, `raw` (values kept as original text), `markers` (unknown/no-value markers), and the reference columns `sources`, `origin_qid`, `origin_absent`, `origin_local`, `origin_edges`, `source_qualifiers`, `ref_mismatch` (must be 0).

In directory mode each run overwrites version 1 of each document; when the content is unchanged no file is touched, and Edge IDs are carried over from the previously published version, so unchanged statements keep their numbers. Pack mode writes a new set, so IDs start from 1 as for a document never published. Lexemes (`L…`, dump 2) are not converted.

## Documentation

Detailed rules (in Korean):

- [docs/conversion.md](docs/conversion.md) — document IDs, record order, value conversion, language codes, property datatypes, sitelinks
- [docs/provenance.md](docs/provenance.md) — document-level sources, TRANSFORMED-BY, conversion of Wikidata references
- [docs/database.md](docs/database.md) — tables and columns read from the wikipg schema

The GEUL grammar itself is specified at [geul.org/grammar](https://geul.org/grammar/).

## Layout

```
cmd/wikigeul/      command: range loop, property documents, name index, report
internal/source/   reading a loaded dump (dump, entities, labels, claims, qualifiers, references, sitelinks)
internal/port/     value conversion, document building, provenance
internal/publish/  document names and the name index
testdata/          database test fixture
```

## Development

```sh
go build ./... && go vet ./... && go test ./...
```

Database tests create a temporary schema through the `PG*` environment variables and are skipped when no database is reachable.

## License

Our own work in this repository is dedicated to the public domain under [CC0 1.0](LICENSE). Third-party components keep their own licenses; see [NOTICE](NOTICE). Wikidata content and the documents wikigeul produces from it are CC0.
