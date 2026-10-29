//ff:func feature=cli type=command control=sequence
//ff:what packWriter.halt — stop a strict run at anomalies: checkpoint what was put, write the anomaly report, return an anomalyError
package main

import "fmt"

// halt checkpoints the stream and report up to the last entity put (the one
// before the anomalous result, which is not written), so -resume starts
// again at that entity, then writes the anomaly report (writeAnomalies) and
// returns an anomalyError (exit 3).
func (w *packWriter) halt(as []anomaly) error {
	if err := w.checkpoint(); err != nil {
		return fmt.Errorf("anomaly at %s, checkpoint failed: %w", as[0].ID, err)
	}
	head := fmt.Sprintf("# wikigeul -strict: stopped at %s; checkpoint after %c%d, %d documents\n%s", as[0].ID, w.prog.Kind, w.prog.Num, w.prog.Docs, w.run)
	if err := writeAnomalies(w.anomaly, head, as); err != nil {
		return fmt.Errorf("anomaly at %s, report not written: %w", as[0].ID, err)
	}
	return anomalyError{ID: as[0].ID, Count: len(as), Path: w.anomaly}
}
