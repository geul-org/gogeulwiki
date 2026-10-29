//ff:type feature=port type=model
//ff:what Report — what one ported document holds and what it left out, with the properties it used
package port

// Report tells what one ported document holds and what it left out (Missing:
// the entity is not in the dump, so no document was written): counts
// of LABEL, DESCRIPTION, ALIAS, TRIPLE and QUALIFIER records, sitelink pages
// and their badges, skipped claims and
// qualifiers by reason (unsupported value type, Lexeme value …), dates
// kept as original text by conversion reason (Raw), somevalue/novalue snaks
// written as slot markers by marker name (Markers; qualifier ones prefixed
// "qualifier ") and the properties the
// document uses as predicates, qualifier properties or values, for porting
// those next. The references (edge/SOURCE, README 출처) count SOURCE Edges,
// their origins by kind (OriginQIDs: the P248 item itself; OriginNone: no
// description, ABSENT; OriginLocals: LOCAL nodes declared, each once
// per document), the TRIPLEs and REFERENCE-NODEs describing those LOCALs
// (OriginEdges), the QUALIFIERs on SOURCE Edges (SourceQuals) and the
// TRIPLEs whose references do not read back from the document as the same
// snak sets (RefMismatch, referenceSnaks; 0 when lossless). Notes holds one
// Note per Skipped and Raw count, in document order, naming the property and
// the original value (noteSkip, noteRaw).
type Report struct {
	ID           string
	Missing      bool
	Labels       int
	Descriptions int
	Aliases      int
	Triples      int
	Qualifiers   int
	Sitelinks    int
	Badges       int
	Skipped      map[string]int
	Raw          map[string]int
	Markers      map[string]int
	Properties   []string
	Sources      int
	OriginQIDs   int
	OriginNone   int
	OriginLocals int
	OriginEdges  int
	SourceQuals  int
	RefMismatch  int
	Notes        []Note
}
