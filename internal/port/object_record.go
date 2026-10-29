//ff:func feature=port type=encoder control=selection
//ff:what objectRecord — the GEUL object record of a claim value (somevalue/novalue slot marker, entity, string or URI, time, monolingual text, coordinate); otherwise the reason the claim is skipped
package port

import (
	"github.com/geul-org/gogeul/pkg/core"

	"github.com/geul-org/wikigeul/internal/source"
)

// objectRecord converts the value of c: an entity ID to its Reference
// (entityRef), a string by its datatype (stringRecord: URI for url and
// Commons pages, else UTF8-TEXT), a time to GREGORIAN-DATE,
// EDTF or GEDCOM-DATE (timeRecord), a monolingualtext to UTF8-LANGUAGE-TEXT
// (monolingualRecord), a globecoordinate to COORDINATE-EARTH or -GEO
// (coordinateRecord), a quantity to QUANTITY (quantityRecord). A time or coordinate that cannot be converted has a
// "time: …" or "globecoordinate: …" reason; claimObject keeps it as original
// text. A somevalue or novalue snak is its slot marker UNKNOWN-VALUE or
// NO-VALUE (snakMarker); an empty value of a value snak is skipped. Other value
// types are not converted in this pass (README 작업 범위); the reason
// names what was skipped.
func objectRecord(c source.Claim) (core.Record, string) {
	if m, ok := snakMarker(c.SnakType); ok {
		return m, ""
	}
	if c.Value == "" {
		return core.Record{}, "빈 값"
	}
	switch c.Type {
	case "wikibase-entityid", "wikibase-item", "wikibase-property":
		if r, ok := entityRef(c.Value); ok {
			return r, ""
		}
		return core.Record{}, "개체: " + c.Value[:1]
	case "string", "external-id":
		return stringRecord(c)
	case "time":
		return timeRecord(c.Value)
	case "monolingualtext":
		return monolingualRecord(c.Value)
	case "globecoordinate", "globe-coordinate":
		return coordinateRecord(c.Value)
	case "quantity":
		return quantityRecord(c.Value)
	default:
		return core.Record{}, "형식 미지원: " + c.Type
	}
}
