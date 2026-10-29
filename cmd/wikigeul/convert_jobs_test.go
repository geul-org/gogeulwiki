//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestConvertJobs — two workers convert every job once and the channel closes
package main

import (
	"database/sql"
	"testing"
)

func TestConvertJobs(t *testing.T) {
	db, schema := fixtureDB(t)
	c := packTestConfig(t, schema)
	jobs := make(chan packJob, 4)
	for i, n := range []uint64{1, 2, 42} {
		jobs <- packJob{Seq: int64(i), Kind: 'Q', Num: n}
	}
	close(jobs)
	seen := map[int64]bool{}
	for r := range convertJobs([]*sql.DB{db, db}, c, packTestProv(t, c), packTestStream(t), jobs) {
		if r.Err != nil || seen[r.Job.Seq] {
			t.Fatalf("%+v", r.Err)
		}
		seen[r.Job.Seq] = true
	}
	if len(seen) != 3 {
		t.Fatal(seen)
	}
}
