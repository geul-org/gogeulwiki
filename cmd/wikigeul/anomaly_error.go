//ff:type feature=cli type=model
//ff:what anomalyError — a strict run stopped at anomalies: the first entity, how many, where the report is
package main

// anomalyError is returned when a strict run stops (packWriter.halt); main
// exits 3 on it.
type anomalyError struct {
	ID    string
	Count int
	Path  string
}
