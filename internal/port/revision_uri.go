//ff:func feature=port type=encoder control=sequence
//ff:what revisionURI — the permanent address of one Wikidata revision (index.php?oldid=)
package port

import "strconv"

// revisionURI is the permanent link of revision revid on Wikidata
// (MediaWiki index.php?oldid=), the same content whatever is edited later.
func revisionURI(revid int64) string {
	return "https://www.wikidata.org/w/index.php?oldid=" + strconv.FormatInt(revid, 10)
}
