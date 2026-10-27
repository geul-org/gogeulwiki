//ff:type feature=port type=model
//ff:what precisionUnits — the units a Wikidata coordinate precision is tried in, in order: degree, arcminute, arcsecond, with degrees per unit
package port

// precisionUnits are the units a globecoordinate precision (degrees) is
// written in, in the order coordinatePrecision tries them: the Wikidata unit
// items degree (Q28390), arcminute (Q209426) and arcsecond (Q829073), with
// how many of the unit make a degree.
var precisionUnits = []struct {
	QID    string
	PerDeg int64
}{
	{"Q28390", 1}, {"Q209426", 60}, {"Q829073", 3600},
}
