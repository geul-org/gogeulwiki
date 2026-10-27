//ff:func feature=port type=encoder control=iteration dimension=1
//ff:what TestCoordinatePrecision — precision texts of the loaded dump (Q1–Q10000 sample) map to degree, arcminute or arcsecond and read back byte for byte; the shortest-form and noisy texts and negatives do not
package port

import (
	"testing"
)

func TestCoordinatePrecision(t *testing.T) {
	type want struct {
		x    float32
		unit uint64
	}
	ok := map[string]want{
		"0.0000010": {1e-6, 28390}, "0.000010": {1e-5, 28390}, "0.0001": {1e-4, 28390}, "0.001": {1e-3, 28390},
		"0.01": {0.01, 28390}, "0.1": {0.1, 28390}, "1": {1, 28390}, "10": {10, 28390}, "0": {0, 28390},
		"0.00000010": {1e-7, 28390}, "0.000000000000010": {1e-14, 28390}, "0.0166667": {0.0166667, 28390},
		"0.000277778": {0.000277778, 28390}, "0.0000277778": {0.0000277778, 28390},
		"0.016666666666667": {1, 209426}, "0.00027777777777778": {1, 829073}, "0.000027777777777778": {0.1, 829073},
		"0.0000027777777777778": {0.01, 829073}, "0.00000027777777777778": {0.001, 829073},
	}
	for p, w := range ok {
		x, unit, good := coordinatePrecision(p)
		if !good || x != w.x || unit.Payload[len(unit.Payload)-1] != uint16Low(w.unit) {
			t.Fatalf("%s: %v %+v %v", p, x, unit, good)
		}
	}
	for _, p := range []string{"0.27155454677066", "0.0002777777777777778", "0.016666666666666666", "0.0010000000000000002", "-0.1", "x"} {
		if _, _, good := coordinatePrecision(p); good {
			t.Fatalf("%s: want no lossless unit", p)
		}
	}
}
