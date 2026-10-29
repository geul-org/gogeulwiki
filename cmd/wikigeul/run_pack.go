//ff:func feature=cli type=command control=iteration dimension=1
//ff:what runPack — pack mode: convert the run's entities with parallel workers and stream them in key order into the Issuer's pack set (zstd), checkpointing and resuming
package main

import (
	"database/sql"
	"database/sql/driver"
	"path/filepath"
	"time"

	"github.com/geul-org/wikigeul/internal/port"
)

// runPack writes the pack set of wikidata.geul.org under
// <root>/.well-known/geul/0x00010000 as one pack stream: the entities of
// -ids, or every property then Q<From>…Q<To> the dump holds (produceJobs),
// are converted by c.Workers goroutines with their own connections
// (convertJobs) and put in document key order (packWriter.take): each
// document as version 1 with its version list, served under h/ by content
// hash without a copy. No documents.geul is written: at Wikidata's size the
// name index would be one 4 GB document, and a document's name follows from
// its ID (port.DocNumber). db (JIT off) lists the entities and builds the
// dictionary samples.
func runPack(c config, conn driver.Connector, db *sql.DB, prov port.Provenance) error {
	if c.RecordedAt == 0 && !c.Resume {
		c.RecordedAt = time.Now().Unix()
	}
	ids, idsHash, err := readIDs(c.IDs)
	if err != nil {
		return err
	}
	set := filepath.Join(c.Root, ".well-known", "geul", "0x00010000")
	s, p, err := openPackStream(set, c, prov, db, ids, idsHash)
	if err != nil {
		return err
	}
	c.RecordedAt = p.RecordedAt
	w, err := newPackWriter(s, c, p)
	if err != nil {
		return err
	}
	defer w.rep.Close()
	jobs := make(chan packJob, 4*c.Workers)
	var perr error
	go produceJobs(db, c, p, ids, jobs, &perr)
	dbs := make([]*sql.DB, c.Workers)
	for i := range dbs {
		dbs[i] = openDB(conn)
	}
	return w.drain(convertJobs(dbs, c, prov, s, jobs), &perr)
}
