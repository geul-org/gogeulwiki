//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what TestNormalReason — Lexeme/EntitySchema skips, validated raw reasons and Warn E6 pass; any other skip, raw reason, Warn or kind does not
package main

import "testing"

func TestNormalReason(t *testing.T) {
	for _, c := range []struct {
		kind, reason string
		want         bool
	}{
		{"skipped", "개체: L", true}, {"skipped", "qualifier 개체: E", true}, {"skipped", "reference 개체: L", true},
		{"skipped", "개체: M", false}, {"skipped", "형식 미지원: x", false}, {"skipped", "언어 코드: x y", false},
		{"raw", "globecoordinate: 정밀도", true}, {"raw", "qualifier time: GEDCOM 대응 없는 정밀도 세기", true},
		{"raw", "time: 0년", false}, {"warn", "E6", true}, {"warn", "E7", false}, {"error", "E6", false},
	} {
		if got := normalReason(c.kind, c.reason); got != c.want {
			t.Errorf("%s %q: %v", c.kind, c.reason, got)
		}
	}
}
