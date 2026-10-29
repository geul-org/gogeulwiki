//ff:func feature=cli type=command control=iteration dimension=1
//ff:what buildAll — the documents of jobs (buildOne), skipping entities the dump lacks
package main

import (
	"database/sql"

	"github.com/geul-org/wikigeul/internal/port"
)

// buildAll builds each job's document in order; entities the dump does not
// have give none. The first error stops it.
func buildAll(db *sql.DB, c config, prov port.Provenance, js []packJob) ([][]byte, error) {
	var out [][]byte
	for _, j := range js {
		b, _, ok, err := buildOne(db, c, prov, j)
		if err != nil {
			return nil, err
		}
		if ok {
			out = append(out, b)
		}
	}
	return out, nil
}
