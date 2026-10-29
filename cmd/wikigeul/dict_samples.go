//ff:func feature=cli type=command control=iteration dimension=1
//ff:what dictSamples — the documents a pack-mode run trains its zstd dictionary on: the first Q entities of the run, built as they will be written
package main

import (
	"database/sql"

	"github.com/geul-org/wikigeul/internal/port"
	"github.com/geul-org/wikigeul/internal/source"
)

// dictSamples builds the documents of the first c.Samples Q entities of the
// run (the ID list's, else the dump's from Q<From>) — items are almost all of
// Wikidata, and the same input always gives the same samples, so the same
// dictionary (pack.StreamOptions.Samples).
func dictSamples(db *sql.DB, c config, prov port.Provenance, ids []packJob) ([][]byte, error) {
	var qs []packJob
	for _, j := range ids {
		if j.Kind == 'Q' && len(qs) < c.Samples {
			qs = append(qs, j)
		}
	}
	if ids == nil {
		nums, err := source.EntityNumbers(db, source.Store{Schema: c.Schema, Dump: c.Dump}, 'Q', c.From-1, c.To, c.Samples)
		if err != nil {
			return nil, err
		}
		qs = jobsOf('Q', nums)
	}
	return buildAll(db, c, prov, qs)
}
