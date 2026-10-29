//ff:func feature=cli type=command control=sequence
//ff:what TestConvertWorker — each job's result goes out and the group is told when the jobs end
package main

import (
	"sync"
	"testing"
)

func TestConvertWorker(t *testing.T) {
	db, schema := fixtureDB(t)
	c := packTestConfig(t, schema)
	jobs, out := make(chan packJob, 1), make(chan packResult, 1)
	jobs <- packJob{Kind: 'Q', Num: 1}
	close(jobs)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	convertWorker(db, c, packTestProv(t, c), packTestStream(t), jobs, out, wg)
	wg.Wait()
	if r := <-out; r.Err != nil || r.Doc != 0x5100000000000001 {
		t.Fatalf("%+v", r.Err)
	}
}
