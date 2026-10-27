//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what TestCoordinateRecord — Berlin becomes the docs' COORDINATE-EARTH (1 arcminute) and its precision text comes back, Mars a COORDINATE-GEO with no precision (unknown); bad JSON, globe, altitude, range and a precision no unit gives back give reasons
package port

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCoordinateRecord(t *testing.T) {
	e := "http://www.wikidata.org/entity/"
	r, reason := coordinateRecord(`{"latitude":52.516666666667,"longitude":13.383333333333,"altitude":null,"precision":0.016666666666667,"globe":"` + e + `Q2"}`)
	ws, _ := core.EncodeRecord(r)
	if reason != "" || !slices.Equal(ws, []core.Word{0x0E84, 0x404A, 0x4222, 0x2222, 0x2251, 0x402A, 0xC444, 0x4444, 0x4389, 0x3F80, 0x0000, 0x0120, 0x0003, 0x3212}) {
		t.Fatalf("berlin: %q %04X", reason, ws)
	}
	if c, _ := core.DecodeCoordinate(r); precisionText(c.Uncertainty, 60) != "0.016666666666667" {
		t.Fatalf("berlin precision back: %q", precisionText(c.Uncertainty, 60))
	}
	r, reason = coordinateRecord(`{"latitude":4.5,"longitude":137.4,"globe":"` + e + `Q111"}`)
	ws, _ = core.EncodeRecord(r)
	if reason != "" || !slices.Equal(ws, []core.Word{0x0E85, 0x00A0, 0x006F, 0x4012, 0, 0, 0, 0x4061, 0x2CCC, 0xCCCC, 0xCCCD, 0, 0, core.Absent}) {
		t.Fatalf("mars: %q %04X", reason, ws)
	}
	r, _ = coordinateRecord(`{"latitude":1,"longitude":2,"precision":0.0000010,"globe":"` + e + `Q405"}`)
	if c, _ := core.DecodeCoordinate(r); c.Globe == nil || c.Uncertainty != float32(1e-6) || c.Unit == nil || c.Unit.Payload[0] != 28390 {
		t.Fatalf("moon, 1e-6 degree kept: %+v", c)
	}
	for want, v := range map[string]string{
		"globecoordinate: JSON 아님": "{",
		"globecoordinate: 천체 없음":   `{"latitude":1,"longitude":2,"precision":0.1,"globe":""}`,
		"globecoordinate: 고도":      `{"latitude":1,"longitude":2,"altitude":120,"precision":0.1,"globe":"` + e + `Q2"}`,
		"globecoordinate: 범위":      `{"latitude":1,"longitude":200,"precision":0.1,"globe":"` + e + `Q2"}`,
		"globecoordinate: 정밀도":     `{"latitude":1,"longitude":2,"precision":-0.00001,"globe":"` + e + `Q2"}`,
	} {
		if _, got := coordinateRecord(v); got != want {
			t.Fatalf("%s: %q want %q", v, got, want)
		}
	}
}
