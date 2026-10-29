//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what bodyRecords — the BODY records of a decoded document, in order
package main

import (
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
)

// bodyRecords returns the records of doc's BODY entries in order, as
// codecspec.Collect reads them.
func bodyRecords(doc inspect.Document) []core.Record {
	body := make([]core.Record, len(doc.Body))
	for i, e := range doc.Body {
		body[i] = e.Record
	}
	return body
}
