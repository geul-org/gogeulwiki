//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what TestAnomalyClass — missing is a data anomaly, errors and checks are other, entity kinds and new raw forms are grammar gaps, unknown types/sites/languages/calendars are porter gaps, malformed values are data anomalies
package main

import "testing"

func TestAnomalyClass(t *testing.T) {
	for _, c := range []struct{ kind, reason, want string }{
		{"missing", "", "데이터 이상"}, {"error", "x", "기타"}, {"check", "F1", "기타"}, {"ref_mismatch", "", "기타"},
		{"skipped", "qualifier 개체: M", "문법 부족"}, {"skipped", "형식 미지원: x", "변환기 미지원"}, {"skipped", "데이터 타입: x", "변환기 미지원"},
		{"skipped", "sitelink 사이트: x", "변환기 미지원"}, {"skipped", "언어 코드: x", "변환기 미지원"}, {"raw", "time: 달력 x", "변환기 미지원"},
		{"skipped", "빈 값", "데이터 이상"}, {"skipped", "rank: odd", "데이터 이상"}, {"skipped", "sitelink badge: P1", "데이터 이상"},
		{"raw", "reference time: JSON 아님", "데이터 이상"}, {"skipped", "문자열: UTF-8 아님", "데이터 이상"}, {"raw", "quantity: 수", "데이터 이상"},
		{"raw", "time: 0년", "데이터 이상"}, {"raw", "globecoordinate: 고도", "문법 부족"}, {"raw", "quantity: 단위", "문법 부족"},
	} {
		if got := anomalyClass(c.kind, c.reason); got != c.want {
			t.Errorf("%s %q: %s, want %s", c.kind, c.reason, got, c.want)
		}
	}
}
