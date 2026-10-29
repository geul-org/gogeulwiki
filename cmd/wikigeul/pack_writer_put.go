//ff:func feature=cli type=command control=iteration dimension=1
//ff:what packWriter.put — store one converted entity: anomaly stop (strict), report line, TRANSFORMED-BY check, version list and version 1 into the stream, checkpoint and progress line when due
package main

import (
	"fmt"
	"time"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/pack"
)

// put writes r's report line and, for a document, puts its version list
// (slot 0) and version 1 (slot 1) under wikidata.geul.org after checking its
// TRANSFORMED-BY records equal the run's first document's (one porter and
// one library commit per set). A strict writer first stops at r's
// anomalies (anomaliesOf, halt) before anything of r is written. Every
// `every` documents it checkpoints and prints a progress line.
func (w *packWriter) put(r packResult) error {
	if w.strict {
		if as := anomaliesOf(r, w.allow); len(as) > 0 {
			return w.halt(as)
		}
	}
	if r.Err != nil {
		return fmt.Errorf("%c%d: %w", r.Job.Kind, r.Job.Num, r.Err)
	}
	n, _ := w.w.WriteString(reportLine(r.Rep) + "\n")
	w.prog.Report += int64(n)
	w.prog.Kind, w.prog.Num = r.Job.Kind, r.Job.Num
	if r.Missing {
		return nil
	}
	tb := transformedBy(r.Raw)
	if w.tb == "" {
		w.tb = tb
	}
	if tb != w.tb {
		return fmt.Errorf("%c%d: TRANSFORMED-BY differs from the run's first document", r.Job.Kind, r.Job.Num)
	}
	for slot, rec := range []pack.StreamRecord{r.List, r.Version} {
		if err := w.s.Put([]core.Word{1, 0}, r.Doc, uint32(slot), rec); err != nil {
			return err
		}
	}
	w.prog.Docs++
	addCounts(w.skipped, r.Rep.Skipped)
	addCounts(w.raw, r.Rep.Raw)
	if w.since++; w.since >= w.every {
		if err := w.checkpoint(); err != nil {
			return err
		}
		w.logProgress(time.Now())
	}
	return nil
}
