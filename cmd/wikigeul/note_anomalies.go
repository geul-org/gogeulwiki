//ff:func feature=cli type=command control=iteration dimension=1
//ff:what noteAnomalies — the skipped and raw Notes of entity id that are neither normal nor allowed, as anomalies
package main

import "github.com/geul-org/wikigeul/internal/port"

// noteAnomalies keeps each Note of id whose reason is not normal
// (normalReason) and not in allow as "kind:reason" (reason without its
// qualifier/reference prefix), classed by anomalyClass.
func noteAnomalies(id string, notes []port.Note, allow map[string]bool) []anomaly {
	var out []anomaly
	for _, n := range notes {
		if normalReason(n.Kind, n.Reason) || allow[n.Kind+":"+bareReason(n.Reason)] {
			continue
		}
		out = append(out, anomaly{ID: id, Kind: n.Kind, Property: n.Property, Value: n.Value, Reason: n.Reason, Class: anomalyClass(n.Kind, n.Reason)})
	}
	return out
}
