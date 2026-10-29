//ff:type feature=cli type=model
//ff:what packResult — one converted entity ready to put: document ID, encoded version list and version, report, or the error
package main

import (
	"github.com/geul-org/gogeul/pkg/pack"

	"github.com/geul-org/wikigeul/internal/port"
)

// packResult is job's output: the document ID, its version list (slot 0) and
// version 1 encoded for the stream, the raw document (for the TRANSFORMED-BY
// check), the report; Missing when the dump lost the entity, Err on failure.
type packResult struct {
	Job     packJob
	Doc     uint64
	List    pack.StreamRecord
	Version pack.StreamRecord
	Raw     []byte
	Rep     port.Report
	Missing bool
	Err     error
}
