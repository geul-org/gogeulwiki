//ff:func feature=cli type=command control=sequence
//ff:what TestFactAnomalies — Warn E6 and allowed Warns pass; other Warns and every Error become anomalies naming the record and place
package main

import (
	"strings"
	"testing"

	"github.com/geul-org/gogeul/pkg/schema"
)

func TestFactAnomalies(t *testing.T) {
	facts := []checkFact{
		{Doc: "version", Fact: schema.Fact{Rule: "E6", Severity: schema.Warn}},
		{Doc: "version", Fact: schema.Fact{Rule: "E9", Severity: schema.Warn}},
		{Doc: "version", Fact: schema.Fact{Rule: "R3", Severity: schema.Warn}},
		{Doc: "list", Fact: schema.Fact{Rule: "E6", Severity: schema.Error, Where: schema.Where{Record: 4, Part: "Object"}, Expected: "e", Actual: "a"}},
	}
	got := factAnomalies("Q1", facts, allowSet([]string{"warn:E9"}))
	if len(got) != 2 || got[0].Reason[:2] != "R3" || got[1].Property != "list record 4 Object" || got[1].Value != "a" || !strings.Contains(got[1].Reason, "expected e") || got[1].Kind != "check" || got[1].Class != "기타" {
		t.Fatalf("%+v", got)
	}
}
