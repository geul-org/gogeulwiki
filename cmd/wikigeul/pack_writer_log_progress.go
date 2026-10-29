//ff:func feature=cli type=command control=sequence
//ff:what packWriter.logProgress — print one progress line after a checkpoint: time, last entity, documents (whole run and this process), rate, skipped and raw reasons of this process
package main

import (
	"fmt"
	"time"
)

// logProgress prints, after a checkpoint, a line a watcher can parse:
// "progress <RFC 3339> last=<entity> docs=<run total> session_docs=<since
// start> rate=<docs/s since start> skipped={…} raw={…}" with the skip and
// raw reasons (reasonCounts) of the documents put since the process
// started.
func (w *packWriter) logProgress(now time.Time) {
	n := w.prog.Docs - w.docs0
	rate := float64(n) / max(now.Sub(w.start).Seconds(), 1e-9)
	fmt.Printf("progress %s last=%c%d docs=%d session_docs=%d rate=%.1f skipped={%s} raw={%s}\n",
		now.Format(time.RFC3339), w.prog.Kind, w.prog.Num, w.prog.Docs, n, rate, reasonCounts(w.skipped), reasonCounts(w.raw))
}
