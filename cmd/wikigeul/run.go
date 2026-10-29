//ff:func feature=cli type=command control=sequence
//ff:what run — port the QID range, then the properties they use, write the name index and the report
package main

import (
	"database/sql"
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/resolve"
	"github.com/geul-org/wikigeul/internal/source"
	"github.com/lib/pq"
)

// run reads the loaded dump's provenance (source.LoadDump, provenanceOf) and
// ports Q<from>…Q<to>, then every property those documents use, all under
// the one Issuer wikidata.geul.org (0x00010000; document IDs by
// port.DocNumber), writes its documents.geul and, when asked, the report.
// A DSN lib/pq cannot parse fails before any work (pq.NewConnector). It
// keeps one connection with JIT off: the planner overestimates the
// reference join (statement_reference ⋈ reference_snak) and JIT-compiles it,
// about 160 ms per entity against 0.5 ms without (2026-10-27). When the
// schema has reference_snak_order (source.Dump.RefOrder) every reference
// occurrence keeps its own snak order; the run says which on stdout.
func run(args []string) error {
	c, err := parseFlags(args)
	if err != nil {
		return err
	}
	conn, err := pq.NewConnector(c.DSN)
	if err != nil {
		return err
	}
	db := sql.OpenDB(conn)
	defer db.Close()
	db.SetMaxOpenConns(1)
	db.Exec("SET jit = off") // speed only: an error leaves JIT on
	d, err := source.LoadDump(db, source.Store{Schema: c.Schema, Dump: c.Dump})
	if err != nil {
		return fmt.Errorf("dump %s/%d: %w", c.Schema, c.Dump, err)
	}
	c.RefOrder = d.RefOrder
	reg, err := registryDir(c.LangRoot)
	if err != nil {
		return err
	}
	if c.Langs, c.LangDoc, err = languageRegistry(reg, c.LangVer); err != nil {
		return err
	}
	fmt.Printf("language registry: %s, %d tags\n", c.LangDoc, len(c.Langs))
	fmt.Printf("reference snak order: per occurrence (reference_snak_order) %v\n", c.RefOrder)
	prov, err := provenanceOf(d, c.Commit, c.LibCommit)
	if err != nil {
		return err
	}
	if c.Out == "pack" {
		return runPack(c, conn, db, prov)
	}
	dir := resolve.Dir{Root: c.Root}
	issuer := []core.Word{1, 0}
	qReps, qDocs, err := portIDs(db, dir, "Q", rangeNumbers(c.From, c.To), issuer, c, prov)
	if err != nil {
		return err
	}
	pReps, pDocs, err := portIDs(db, dir, "P", propertyNumbers(qReps), issuer, c, prov)
	if err != nil {
		return err
	}
	if err := writeIndex(c.Root, issuer, append(qDocs, pDocs...)); err != nil {
		return err
	}
	fmt.Printf("Q: %d documents of %d IDs, P: %d documents of %d IDs\n", len(qDocs), len(qReps), len(pDocs), len(pReps))
	if c.Report == "" {
		return nil
	}
	return writeReport(c.Report, runLine(c), append(qReps, pReps...))
}
