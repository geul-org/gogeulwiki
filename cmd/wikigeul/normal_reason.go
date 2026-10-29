//ff:func feature=cli type=helper control=selection
//ff:what normalReason — whether a skip, raw-kept value or schema Warn is one of the normal ones a strict run lets pass
package main

// normalReason says whether a Note or Fact of kind is normal (README 엄격
// 모드), its reason taken without the qualifier/reference prefix:
//   - skipped: only values of a Lexeme (L, its forms and senses) or an
//     EntitySchema (E), for which GEUL has no Reference Codec yet;
//   - raw: the original-text cases of docs/conversion.md met in the
//     validated runs (normalRaw);
//   - warn: E6 (NO-VALUE beside a value, a property of the source data).
//
// Anything else — another skip, a new raw reason, any Error Fact — is an
// anomaly.
func normalReason(kind, reason string) bool {
	r := bareReason(reason)
	switch kind {
	case "skipped":
		return r == "개체: L" || r == "개체: E"
	case "raw":
		return normalRaw(r)
	case "warn":
		return r == "E6"
	default:
		return false
	}
}
