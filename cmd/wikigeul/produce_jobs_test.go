//ff:func feature=cli type=command control=sequence
//ff:what TestProduceJobs — every property then the Q range in order and numbered; from a checkpoint in P or in Q; the ID list after the checkpoint; a bad schema leaves an error
package main

import (
	"testing"
)

func TestProduceJobs(t *testing.T) {
	db, schema := fixtureDB(t)
	c := config{Schema: schema, Dump: 1, From: 2, To: 40}
	collect := func(c config, p packProgress, ids []packJob) ([]string, error) {
		jobs := make(chan packJob, 100)
		var err error
		produceJobs(db, c, p, ids, jobs, &err)
		var out []string
		for j := range jobs {
			out = append(out, string(j.Kind)+string(rune('0'+j.Seq)))
		}
		return out, err
	}
	all, err := collect(c, packProgress{}, nil)
	if err != nil || len(all) != 6 || all[0] != "P0" || all[2] != "Q2" || all[5] != "Q5" {
		t.Fatalf("%v %v", all, err)
	}
	if got, _ := collect(c, packProgress{Kind: 'P', Num: 31}, nil); len(got) != 5 || got[0] != "P0" {
		t.Fatalf("after P31: %v", got)
	}
	if got, _ := collect(c, packProgress{Kind: 'Q', Num: 3}, nil); len(got) != 2 {
		t.Fatalf("after Q3: %v", got)
	}
	if got, _ := collect(c, packProgress{Kind: 'Q', Num: 1}, []packJob{{Kind: 'P', Num: 1}, {Kind: 'Q', Num: 1}, {Kind: 'Q', Num: 9}}); len(got) != 1 {
		t.Fatalf("ID list: %v", got)
	}
	if _, err := collect(config{Schema: "1x", Dump: 1}, packProgress{}, nil); err == nil {
		t.Fatal("bad schema: want error")
	}
}
