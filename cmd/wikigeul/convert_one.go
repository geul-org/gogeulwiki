//ff:func feature=cli type=command control=sequence
//ff:what convertOne — one entity as a packResult: its document and first version list (fixed RECORDED-AT), both encoded for the stream
package main

import (
	"database/sql"

	"github.com/geul-org/gogeul/pkg/pack"
	"github.com/geul-org/gogeul/pkg/resolve"

	"github.com/geul-org/wikigeul/internal/port"
)

// convertOne builds job's document (buildOne), its version list as
// directory mode's Overwrite would write it but recorded at c.RecordedAt
// (resolve.FirstVersionList), and encodes both (s.Encode, safe on many
// goroutines).
func convertOne(db *sql.DB, c config, prov port.Provenance, s *pack.Stream, job packJob) packResult {
	r := packResult{Job: job}
	b, rep, ok, err := buildOne(db, c, prov, job)
	r.Rep, r.Missing, r.Err = rep, !ok && err == nil, err
	if !ok {
		return r
	}
	r.Doc, _ = port.DocNumber(job.Kind, job.Num) // BuildDoc accepted the number
	list, err := resolve.FirstVersionList(b, c.RecordedAt)
	if err == nil {
		r.List, err = s.Encode(list)
	}
	if err == nil {
		r.Version, err = s.Encode(b)
	}
	r.Raw, r.Err = b, err
	return r
}
