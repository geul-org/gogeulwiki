//ff:type feature=port type=model
//ff:what Provenance — the source and the tools every ported document records in its HEAD
package port

// Provenance is what every ported document records in its HEAD (README
// 출처 기록): the dump URI, the SHA-256 of the file in hex (empty when
// unknown) and when the dump file was received (POSIX seconds) — the HEAD
// SOURCE on the dump and its QUALIFIERs (dumpSource) — and the tools that
// turned it into GEUL, one TRANSFORMED-BY each in order — wikigeul and the
// GEUL library it builds with (the loader is not one: it stores the dump
// losslessly).
type Provenance struct {
	DumpURI    string
	DumpSHA256 string
	AcquiredAt int64
	Tools      []Tool
}
