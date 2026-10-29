//ff:func feature=cli type=command control=sequence
//ff:what buildOne — load one entity from the dump and build its document with fresh LOCAL IDs (pack mode); false when the dump lacks it
package main

import (
	"database/sql"
	"strconv"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/delta"

	"github.com/geul-org/wikigeul/internal/port"
	"github.com/geul-org/wikigeul/internal/source"
)

// buildOne loads entity job (source.LoadEntity) and builds its document
// (port.BuildDoc) with prov and the run's registry. A pack-mode run writes a
// new set, so no earlier version lends its IDs: they start from 1, as for a
// document never published in directory mode.
func buildOne(db *sql.DB, c config, prov port.Provenance, job packJob) ([]byte, port.Report, bool, error) {
	id := string(job.Kind) + strconv.FormatUint(job.Num, 10)
	e, ok, err := source.LoadEntity(db, source.Store{Schema: c.Schema, Dump: c.Dump, RefOrder: c.RefOrder}, id)
	if err != nil || !ok {
		return nil, port.Report{ID: id, Missing: err == nil}, false, err
	}
	ids, _ := delta.NewLocalIDs(nil, 0) // no previous version: never fails
	b, rep, err := port.BuildDoc(e, []core.Word{1, 0}, prov, ids, c.Langs)
	return b, rep, err == nil, err
}
