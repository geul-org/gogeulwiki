//ff:func feature=port type=helper control=sequence
//ff:what TestCountMarker — markers counted by name and prefix, the map created on first use; other records not counted
package port

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestCountMarker(t *testing.T) {
	var rep Report
	unknown, _ := snakMarker("somevalue")
	none, _ := snakMarker("novalue")
	countMarker(&rep, "qualifier ", unknown)
	countMarker(&rep, "", none)
	countMarker(&rep, "", none)
	countMarker(&rep, "", core.EncodeLocalRef(1))
	if len(rep.Markers) != 2 || rep.Markers["qualifier UNKNOWN-VALUE"] != 1 || rep.Markers["NO-VALUE"] != 2 {
		t.Fatalf("%v", rep.Markers)
	}
	var empty Report
	countMarker(&empty, "", core.EncodeLocalRef(1))
	if empty.Markers != nil {
		t.Fatal("no marker: map not created")
	}
}
