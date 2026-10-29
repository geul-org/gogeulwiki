//ff:func feature=cli type=command control=iteration dimension=1
//ff:what factAnomalies — the schema Facts of entity id but the normal or allowed Warns, as anomalies
package main

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/schema"
)

// factAnomalies keeps every Error Fact and each Warn whose rule is neither
// normal (normalReason "warn": E6) nor allowed ("warn:<rule>"), with the
// record it is on (version, list), its place and what was expected.
func factAnomalies(id string, facts []checkFact, allow map[string]bool) []anomaly {
	var out []anomaly
	for _, f := range facts {
		if f.Severity == schema.Warn && (normalReason("warn", f.Rule) || allow["warn:"+f.Rule]) {
			continue
		}
		where := fmt.Sprintf("%s record %d %s", f.Doc, f.Where.Record, f.Where.Part)
		out = append(out, anomaly{ID: id, Kind: "check", Property: where, Value: f.Actual, Reason: fmt.Sprintf("%s (severity %d): expected %s", f.Rule, f.Severity, f.Expected), Class: "기타"})
	}
	return out
}
