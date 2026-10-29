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

## Usage

```sh
wikigeul -root ./archive -from 1 -to 1000 \
  -commit "$(git rev-parse HEAD)" \
  -lib-commit <gogeul commit> \
  -report report.tsv
```

This converts Q1–Q1000, then every property those documents use, and writes the Issuer's name index `documents.geul` (`q42`, `p31`, …). Documents are written to `<root>/.well-known/geul/`. The index keeps documents already published there, so a range can be converted in several runs.

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

The report starts with a comment line recording the run (`# wikigeul -from … -to … -schema … -dump-id … -commit … -lib-commit …`), then a header and one line per ID in the range. `status` is `ok` (document written) or `missing` (no such entity in the dump). The other columns count what was written and what was skipped: `labels`, `descriptions`, `aliases`, `triples`, `qualifiers`, `sitelinks`, `badges`, `skipped`, `raw` (values kept as original text), `markers` (unknown/no-value markers), and the reference columns `sources`, `origin_qid`, `origin_absent`, `origin_local`, `origin_edges`, `source_qualifiers`, `ref_mismatch` (must be 0).

Each run overwrites version 1 of each document; when the content is unchanged no file is touched. Edge IDs are carried over from the previously published version, so unchanged statements keep their numbers.

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
