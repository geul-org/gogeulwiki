//ff:func feature=cli type=command control=iteration dimension=1
//ff:what convertJobs — start one worker per connection over the jobs; results arrive in any order and the channel closes when all are done
package main

import (
	"database/sql"
	"sync"

	"github.com/geul-org/gogeul/pkg/pack"

	"github.com/geul-org/wikigeul/internal/port"
)

// convertJobs runs convertWorker on each of dbs. The writer restores the
// output order from the job numbers (packWriter.take).
func convertJobs(dbs []*sql.DB, c config, prov port.Provenance, s *pack.Stream, jobs <-chan packJob) <-chan packResult {
	out := make(chan packResult, 4*len(dbs))
	wg := &sync.WaitGroup{}
	for _, db := range dbs {
		wg.Add(1)
		go convertWorker(db, c, prov, s, jobs, out, wg)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
