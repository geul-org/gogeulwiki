//ff:func feature=cli type=helper control=sequence
//ff:what anomalyError.Error — the stop message: entity, count, report path
package main

import "fmt"

// Error says where the strict run stopped and where the report is.
func (e anomalyError) Error() string {
	return fmt.Sprintf("wikigeul -strict: %d anomalies at %s, stopped after a checkpoint; report %s", e.Count, e.ID, e.Path)
}
