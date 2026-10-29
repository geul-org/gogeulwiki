//ff:func feature=cli type=command control=iteration dimension=1
//ff:what packWriter.take — hold a result until every earlier job is put, then put the run of results that are ready
package main

// take keeps r until the results before it have been put, then puts every
// waiting result in job order.
func (w *packWriter) take(r packResult) error {
	w.pending[r.Job.Seq] = r
	for {
		r, ok := w.pending[w.next]
		if !ok {
			return nil
		}
		delete(w.pending, w.next)
		w.next++
		if err := w.put(r); err != nil {
			return err
		}
	}
}
