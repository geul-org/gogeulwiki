//ff:func feature=cli type=helper control=selection
//ff:what normalRaw — whether a raw-kept reason is a normal one (coordinate precision and range, timezone/bounds, coarse precisions, BCE decade, Y-year, day form, year 0, day out of range, the surveyed calendar typos)
//ff:why 2026-10-30 user decision: data errors B–F of the pre-survey are kept as original text and pass; calendar typos are listed one by one so an unseen calendar model still stops the run
package main

import "strings"

// normalRaw lists the raw (original text) reasons a strict run lets pass.
// The validated runs (files/wikidata full-validation-2026-10-29,
// pack-mode-2026-10-30) met a coordinate precision no unit gives back, a
// time with a timezone or bounds, a time precision coarser than a decade
// (0–5, 천년, 세기; and 10년 for Julian dates) that EDTF or GEDCOM has no
// form for, a BCE or 5+ digit decade, a Y-prefixed year go-edtf cannot
// check and a day-precision date that is no calendar day. The pre-survey of
// dump 1 (pre-survey-2026-10-30) added data errors kept as written: a
// Gregorian year 0 below day precision, a day out of the int32 day range, a
// coordinate out of its degree range (Earth longitude beyond ±180, latitude
// beyond ±90) and the six calendar model typos it found, each by its exact
// IRI. Every other raw reason is new.
func normalRaw(r string) bool {
	switch r {
	case "globecoordinate: 정밀도", "globecoordinate: 범위", "time: 시간대·오차 범위", "time: 10년(기원전·큰 연도)",
		"time: EDTF 검증 불가(Y 접두 연도)", "time: 날짜 형식", "time: 0년", "time: 범위",
		"time: 달력 http://www.wikidata.org/entity/Q198572", "time: 달력 http://www.wikidata.org/entity/Q19857227",
		"time: 달력 http://www.wikidata.org/entity/Q1985726", "time: 달력 http://www.wikidata.org/entity/Q1985728",
		"time: 달력 http://www.wikidata.org/entity/Q199322", "time: 달력 http://www.wikidata.org/entity/Q199382":
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
