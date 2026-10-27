//ff:type feature=port type=model
//ff:what originProperties — the reference properties that describe the cited source itself (원천 쪽); every other one describes the citation (인용 쪽)
package port

// originProperties are the Wikidata reference properties that say something
// about the cited work itself (README 출처, files/geul/proposals/record-source.md
// 속성 분류): its address, title, authors, publication, publisher, language,
// and P248 (stated in) when the source is a LOCAL node. A reference snak of
// any other property (P813 retrieved, P143 imported from, P887, P4656, page,
// section, external identifiers, archive URL …) is about the act of citing
// and becomes a QUALIFIER on the SOURCE Edge: an unclassified property is not
// read as a claim about the source.
var originProperties = map[string]bool{
	"P248":  true, // stated in
	"P854":  true, // reference URL
	"P1476": true, // title
	"P1680": true, // subtitle
	"P50":   true, // author
	"P2093": true, // author name string
	"P98":   true, // editor
	"P655":  true, // translator
	"P577":  true, // publication date
	"P1433": true, // published in
	"P123":  true, // publisher
	"P291":  true, // place of publication
	"P407":  true, // language of work or name
}
