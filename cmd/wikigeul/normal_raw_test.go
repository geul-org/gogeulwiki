//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what TestNormalRaw — the validated and pre-surveyed raw reasons pass; finer precisions, other prefixes, other calendars and malformed values do not
package main

import "testing"

func TestNormalRaw(t *testing.T) {
	ok := []string{"globecoordinate: 정밀도", "time: 시간대·오차 범위", "time: 10년(기원전·큰 연도)", "time: EDTF 검증 불가(Y 접두 연도)", "time: 날짜 형식",
		"time: GEDCOM 대응 없는 정밀도 0", "time: GEDCOM 대응 없는 정밀도 10년", "time: EDTF 대응 미결정 정밀도 세기", "time: EDTF 대응 미결정 정밀도 5", "time: EDTF 대응 미결정 정밀도 천년",
		"time: 0년", "time: 범위", "globecoordinate: 범위",
		"time: 달력 http://www.wikidata.org/entity/Q198572", "time: 달력 http://www.wikidata.org/entity/Q19857227",
		"time: 달력 http://www.wikidata.org/entity/Q1985726", "time: 달력 http://www.wikidata.org/entity/Q1985728",
		"time: 달력 http://www.wikidata.org/entity/Q199322", "time: 달력 http://www.wikidata.org/entity/Q199382"}
	bad := []string{"time: GEDCOM 대응 없는 정밀도 12", "time: EDTF 대응 미결정 정밀도 월", "time: JSON 아님", "quantity: 단위", "globecoordinate: 고도", "globecoordinate: 천체 없음", "url: 스킴 없음",
		"time: 달력 x", "time: 달력 http://www.wikidata.org/entity/Q12138", "time: 달력 http://www.wikidata.org/entity/Q1985727x", "time: 달력 ", "time: EDTF 검증 실패", "time: GEDCOM 검증 실패", "monolingualtext: 언어 코드 x", "세기", ""}
	for _, r := range ok {
		if !normalRaw(r) {
			t.Errorf("%q: want normal", r)
		}
	}
	for _, r := range bad {
		if normalRaw(r) {
			t.Errorf("%q: want anomaly", r)
		}
	}
}
