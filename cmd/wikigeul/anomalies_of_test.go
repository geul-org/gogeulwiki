//ff:func feature=cli type=command control=sequence
//ff:what TestAnomaliesOf — an error or a missing entity alone (ID from the job when the report has none); otherwise abnormal notes, a reference mismatch and abnormal Facts; a clean result has none
package main

import (
	"errors"
	"testing"

	"github.com/geul-org/gogeul/pkg/schema"

	"github.com/geul-org/wikigeul/internal/port"
)

func TestAnomaliesOf(t *testing.T) {
	if got := anomaliesOf(packResult{Job: packJob{Kind: 'Q', Num: 5}, Err: errors.New("boom")}, nil); len(got) != 1 || got[0].ID != "Q5" || got[0].Kind != "error" || got[0].Reason != "boom" {
		t.Fatalf("error: %+v", got)
	}
	if got := anomaliesOf(packResult{Rep: port.Report{ID: "Q6"}, Missing: true}, nil); len(got) != 1 || got[0].ID != "Q6" || got[0].Class != "데이터 이상" {
		t.Fatalf("missing: %+v", got)
	}
	r := packResult{Rep: port.Report{ID: "Q7", RefMismatch: 2, Notes: []port.Note{{Kind: "skipped", Reason: "빈 값"}, {Kind: "skipped", Reason: "개체: L"}}},
		Facts: []checkFact{{Doc: "version", Fact: schema.Fact{Rule: "E6", Severity: schema.Warn}}, {Doc: "version", Fact: schema.Fact{Rule: "F1"}}}}
	got := anomaliesOf(r, nil)
	if len(got) != 3 || got[0].Reason != "빈 값" || got[1].Kind != "ref_mismatch" || got[2].Kind != "check" {
		t.Fatalf("%+v", got)
	}
	if got := anomaliesOf(packResult{Rep: port.Report{ID: "Q8"}}, nil); len(got) != 0 {
		t.Fatalf("clean: %+v", got)
	}
}
