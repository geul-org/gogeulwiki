//ff:func feature=cli type=command control=iteration dimension=1
//ff:what packWriter.put — store one converted entity: report line, TRANSFORMED-BY check, version list and version 1 into the stream, checkpoint when due
package main

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/pack"
)

// put writes r's report line and, for a document, puts its version list
// (slot 0) and version 1 (slot 1) under wikidata.geul.org after checking its
// TRANSFORMED-BY records equal the run's first document's (one porter and
// one library commit per set). Every `every` documents it checkpoints.
func (w *packWriter) put(r packResult) error {
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
	if w.since++; w.since >= w.every {
		return w.checkpoint()
	}
	return nil
}
