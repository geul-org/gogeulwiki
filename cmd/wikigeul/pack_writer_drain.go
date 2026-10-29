//ff:func feature=cli type=command control=iteration dimension=1
//ff:what packWriter.drain — put every result in job order, then, if the producer did not fail, close the stream
package main

import "time"

// drain takes every result (take) until the workers are done; a failed put
// stops it. A producer error (*perr, set before jobs closed, so before
// results closed) fails the run instead of closing an incomplete set;
// otherwise the stream is finished.
func (w *packWriter) drain(results <-chan packResult, perr *error) error {
	for r := range results {
		if err := w.take(r); err != nil {
			return err
		}
	}
	if *perr != nil {
		return *perr
	}
	return w.finish(time.Now())
}
