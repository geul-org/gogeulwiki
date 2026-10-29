//ff:func feature=cli type=helper control=selection
//ff:what normalRaw — whether a raw-kept reason is one met and validated before (coordinate precision, timezone/bounds, coarse precisions, BCE decade, Y-year, day form)
package main

import "strings"

// normalRaw lists the raw (original text) reasons the validated runs met
// (files/wikidata full-validation-2026-10-29, pack-mode-2026-10-30): a
// coordinate precision no unit gives back, a time with a timezone or
// bounds, a time precision coarser than a decade (0–5, 천년, 세기; and 10년
// for Julian dates) that EDTF or GEDCOM has no form for, a BCE or 5+ digit
// decade, a Y-prefixed year go-edtf cannot check and a day-precision date
// that is no calendar day. Every other raw reason is new.
func normalRaw(r string) bool {
	switch r {
	case "globecoordinate: 정밀도", "time: 시간대·오차 범위", "time: 10년(기원전·큰 연도)",
		"time: EDTF 검증 불가(Y 접두 연도)", "time: 날짜 형식":
		return true
	}
	p, ok := strings.CutPrefix(r, "time: GEDCOM 대응 없는 정밀도 ")
	if !ok {
		p, ok = strings.CutPrefix(r, "time: EDTF 대응 미결정 정밀도 ")
	}
	switch p {
	case "0", "1", "2", "3", "4", "5", "천년", "세기", "10년":
		return ok
	default:
		return false
	}
}
