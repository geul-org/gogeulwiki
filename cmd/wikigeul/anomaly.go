//ff:type feature=cli type=model
//ff:what anomaly — one thing a strict run stops at: entity, kind, property, original value, reason and its class
package main

// anomaly is one finding that stops a strict run (anomaliesOf): the entity
// ID, the kind (skipped, raw, error, missing, ref_mismatch, check), the
// property it sits on, the original value, the reason as the report counts
// it and the class (anomalyClass: 문법 부족, 변환기 미지원, 데이터 이상,
// 기타).
type anomaly struct {
	ID       string
	Kind     string
	Property string
	Value    string
	Reason   string
	Class    string
}
