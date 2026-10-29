//ff:func feature=cli type=command control=sequence
//ff:what newPackWriter — the pack writer over a stream: the report opened fresh (run line, header) or cut back to the checkpoint's size
package main

import (
	"bufio"
	"os"
	"time"

	"github.com/geul-org/gogeul/pkg/pack"
)

// newPackWriter opens the report c.Report (os.DevNull when none): a new run
// writes the run line and header, a resumed one cuts the file back to
// p.Report and appends (O_APPEND). The strict settings, the run line and
// the progress counters come from c and p.
func newPackWriter(s *pack.Stream, c config, p packProgress) (*packWriter, error) {
	path := c.Report
	if path == "" {
		path = os.DevNull
	}
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o644)
	if err != nil {
		return nil, err
	}
	w := &packWriter{s: s, rep: f, w: bufio.NewWriterSize(f, 1<<20), prog: p, every: c.Every, pending: map[int64]packResult{},
		strict: c.Strict, allow: allowSet(c.Allow), anomaly: c.Anomaly, run: runLine(c), skipped: map[string]int{}, raw: map[string]int{}, start: time.Now(), docs0: p.Docs}
	fresh := p.Docs == 0 && p.Kind == 0
	if c.Report != "" {
		err = f.Truncate(map[bool]int64{true: 0, false: p.Report}[fresh])
	}
	if err != nil {
		f.Close()
		return nil, err
	}
	if fresh {
		head := runLine(c) + "\n" + reportHeader + "\n"
		w.prog.Report = int64(len(head))
		w.w.WriteString(head) // buffered: errors show at Flush
	}
	return w, nil
}
