//ff:func feature=cli type=command control=sequence
//ff:what anomaliesOf — what in one result stops a strict run: an error or panic, a missing entity, notes outside the normal list, reference mismatches, schema Facts other than Warn E6
package main

import (
	"fmt"
	"strconv"
)

// anomaliesOf lists the anomalies of r (README 엄격 모드): a failed
// conversion (r.Err, panics included) or an entity the dump lists but does
// not hold (r.Missing) alone; otherwise each skipped or raw Note whose
// reason is not normal (normalReason) and not allowed (allow
// "kind:reason", reason without its qualifier/reference prefix), a non-zero
// RefMismatch and each schema Fact but an allowed Warn (E6, or "warn:<rule>"
// in allow).
func anomaliesOf(r packResult, allow map[string]bool) []anomaly {
	id := r.Rep.ID
	if id == "" {
		id = string(r.Job.Kind) + strconv.FormatUint(r.Job.Num, 10)
	}
	if r.Err != nil {
		return []anomaly{{ID: id, Kind: "error", Reason: r.Err.Error(), Class: "기타"}}
	}
	if r.Missing {
		return []anomaly{{ID: id, Kind: "missing", Reason: "덤프 목록에 있으나 개체를 읽지 못함", Class: anomalyClass("missing", "")}}
	}
	out := noteAnomalies(id, r.Rep.Notes, allow)
	if r.Rep.RefMismatch > 0 {
		out = append(out, anomaly{ID: id, Kind: "ref_mismatch", Reason: fmt.Sprintf("출처 역변환 불일치 TRIPLE %d개", r.Rep.RefMismatch), Class: "기타"})
	}
	return append(out, factAnomalies(id, r.Facts, allow)...)
}
