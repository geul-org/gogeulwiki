//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what TestNormalRaw — the validated raw reasons and precisions pass; finer precisions, other prefixes and malformed values do not
package main

import "testing"

func TestNormalRaw(t *testing.T) {
	ok := []string{"globecoordinate: 정밀도", "time: 시간대·오차 범위", "time: 10년(기원전·큰 연도)", "time: EDTF 검증 불가(Y 접두 연도)", "time: 날짜 형식",
		"time: GEDCOM 대응 없는 정밀도 0", "time: GEDCOM 대응 없는 정밀도 10년", "time: EDTF 대응 미결정 정밀도 세기", "time: EDTF 대응 미결정 정밀도 5", "time: EDTF 대응 미결정 정밀도 천년"}
	bad := []string{"time: GEDCOM 대응 없는 정밀도 12", "time: EDTF 대응 미결정 정밀도 월", "time: JSON 아님", "time: 0년", "quantity: 단위", "globecoordinate: 고도", "url: 스킴 없음", "time: 달력 x", "세기", ""}
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
