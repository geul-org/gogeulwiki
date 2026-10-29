//ff:func feature=cli type=command control=iteration dimension=1
//ff:what convertWorker — convert jobs until the channel closes, each result to out
package main

import (
	"database/sql"
	"sync"

	"github.com/geul-org/gogeul/pkg/pack"

	"github.com/geul-org/wikigeul/internal/port"
)

// convertWorker runs convertOne on its own connection db for each job and
// signals wg when jobs is closed and drained.
func convertWorker(db *sql.DB, c config, prov port.Provenance, s *pack.Stream, jobs <-chan packJob, out chan<- packResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		out <- convertOne(db, c, prov, s, j)
	}
}
