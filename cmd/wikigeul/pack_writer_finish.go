//ff:func feature=cli type=command control=sequence
//ff:what packWriter.finish — flush the report and close the stream (manifest, latest.txt), printing what was written
package main

import (
	"fmt"
	"time"
)

// finish flushes the report and closes the stream, then prints the
// generation, documents, records and bytes.
func (w *packWriter) finish(start time.Time) error {
	if err := w.w.Flush(); err != nil {
		return err
	}
	res, err := w.s.Close()
	if err != nil {
		return err
	}
	fmt.Printf("pack: generation %d manifest %s, %d documents, %d records (%d zstd), %d objects, %d bytes\n",
		res.Generation, res.Manifest, w.prog.Docs, res.Records, res.Compressed, len(res.Objects), res.Bytes)
	return nil
}
