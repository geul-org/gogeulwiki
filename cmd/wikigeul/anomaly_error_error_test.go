//ff:func feature=cli type=helper control=sequence
//ff:what TestAnomalyErrorError — the message names the count, entity and report path
package main

import "testing"

func TestAnomalyErrorError(t *testing.T) {
	if got := (anomalyError{ID: "Q7", Count: 2, Path: "a.tsv"}).Error(); got != "wikigeul -strict: 2 anomalies at Q7, stopped after a checkpoint; report a.tsv" {
		t.Fatal(got)
	}
}
