//ff:func feature=cli type=helper control=selection
//ff:what jobAfter — whether entity (kind, num) comes after (afterKind, afterNum) in document key order: every P before every Q, then by number
package main

// jobAfter orders entities as their documents' keys do (P = 0x50… before
// Q = 0x51…, then the number). afterKind 0 is before everything.
func jobAfter(kind byte, num uint64, afterKind byte, afterNum uint64) bool {
	switch {
	case afterKind == 0:
		return true
	case kind != afterKind:
		return kind > afterKind
	default:
		return num > afterNum
	}
}
