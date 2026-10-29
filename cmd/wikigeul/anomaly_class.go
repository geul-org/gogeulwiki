//ff:func feature=cli type=helper control=selection
//ff:what anomalyClass — a first guess at why an anomaly happened: 문법 부족, 변환기 미지원, 데이터 이상 or 기타
package main

import "strings"

// anomalyClass classes an anomaly of kind with reason (without its
// qualifier/reference prefix) as a first guess for the reader:
//   - 문법 부족: GEUL has no form for it (an entity kind with no Reference
//     Codec; a value kept as original text for a reason not met before);
//   - 변환기 미지원: the porter does not know it (datavalue type, property
//     datatype, sitelink site, language code);
//   - 데이터 이상: the source is malformed (empty value, bad property ID or
//     rank, non-item badge, not JSON or UTF-8, impossible date or number,
//     an entity the dump lists but cannot load);
//   - 기타: conversion errors and panics, reference mismatches, schema Facts.
func anomalyClass(kind, reason string) string {
	r := bareReason(reason)
	malformed := []string{"빈 값", "속성 ID", "rank:", "badge:", "JSON 아님", "UTF-8 아님", "날짜 형식", "범위", "검증 실패", "0년", "quantity: 수", "quantity: 경계"}
	switch {
	case kind == "missing":
		return "데이터 이상"
	case kind != "skipped" && kind != "raw":
		return "기타"
	case strings.HasPrefix(r, "개체: "):
		return "문법 부족"
	case strings.HasPrefix(r, "형식 미지원: ") || strings.HasPrefix(r, "데이터 타입: ") || strings.HasPrefix(r, "sitelink 사이트: ") || strings.HasPrefix(r, "언어 코드: ") || strings.HasPrefix(r, "time: 달력 "):
		return "변환기 미지원"
	case containsAny(r, malformed):
		return "데이터 이상"
	default:
		return "문법 부족"
	}
}
