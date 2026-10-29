//ff:func feature=port type=encoder control=selection
//ff:what coordinateRecord — a Wikidata globecoordinate as COORDINATE-EARTH or COORDINATE-GEO with its precision as an uncertainty in degrees, arcminutes or arcseconds; otherwise the reason it is not converted
package port

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
)

// coordinateRecord converts a globecoordinate datavalue: on Earth (globe Q2)
// to COORDINATE-EARTH, elsewhere to COORDINATE-GEO with the globe as a
// WIKIDATA-QID. The precision is kept without loss as the uncertainty and its
// unit item (coordinatePrecision, the precision's text as the value has
// it); a missing precision is an unknown uncertainty (unit ABSENT). A value
// with no or a non-item globe, an altitude (the formats hold none), degrees
// out of range (longitude [-180, 180] on Earth, [-360, 360] elsewhere as
// Wikidata allows) or a precision no unit gives back gives its reason
// (claimObject keeps such values as original text).
func coordinateRecord(value string) (core.Record, string) {
	var v struct {
		Lat       float64      `json:"latitude"`
		Lon       float64      `json:"longitude"`
		Precision *json.Number `json:"precision"`
		Altitude  *float64     `json:"altitude"`
		Globe     string       `json:"globe"`
	}
	dec := json.NewDecoder(bytes.NewReader([]byte(value)))
	dec.UseNumber()
	if err := dec.Decode(&v); err != nil {
		return core.Record{}, "globecoordinate: JSON 아님"
	}
	globe, okGlobe := entityRef(strings.TrimPrefix(v.Globe, "http://www.wikidata.org/entity/"))
	earth := v.Globe == "http://www.wikidata.org/entity/Q2"
	lonMax := 360.0
	if earth {
		lonMax = 180
	}
	lonMin := -lonMax
	c := core.Coordinate{Lat: v.Lat, Lon: v.Lon}
	if !earth {
		c.Globe = &globe
	}
	var unit core.Record
	okPrecision := v.Precision == nil
	if v.Precision != nil {
		c.Uncertainty, unit, okPrecision = coordinatePrecision(v.Precision.String())
		c.Unit = &unit
	}
	switch {
	case !okGlobe || globe.Codec != core.RefWikidataQID<<3:
		return core.Record{}, "globecoordinate: 천체 없음"
	case v.Altitude != nil:
		return core.Record{}, "globecoordinate: 고도"
	case v.Lat < -90 || v.Lat > 90 || v.Lon < lonMin || v.Lon > lonMax:
		return core.Record{}, "globecoordinate: 범위"
	case !okPrecision:
		return core.Record{}, "globecoordinate: 정밀도"
	default:
		return core.EncodeCoordinate(c), ""
	}
}
