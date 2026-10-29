//ff:func feature=cli type=command control=sequence
//ff:what packWriter.checkpoint — flush and sync the report, then checkpoint the stream with the progress (JSON) in the same atomic save
package main

import (
	"encoding/json"
	"os"
)

// checkpoint makes the report durable up to prog.Report (os.DevNull, no
// report, does not sync) and saves prog with
// the stream's checkpoint (pack.Stream.Checkpoint), so a resume finds the
// stream, the progress and the report in agreement.
func (w *packWriter) checkpoint() error {
	if err := w.w.Flush(); err != nil {
		return err
	}
	if err := w.rep.Sync(); err != nil && w.rep.Name() != os.DevNull {
		return err
	}
	meta, _ := json.Marshal(w.prog) // plain fields: never fails
	w.since = 0
	return w.s.Checkpoint(meta)
}
