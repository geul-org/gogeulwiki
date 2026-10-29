//ff:type feature=cli type=model
//ff:what packWriter — the pack-mode writer: stream, report file, progress, the expected TRANSFORMED-BY and the results waiting for their turn
package main

import (
	"bufio"
	"os"

	"github.com/geul-org/gogeul/pkg/pack"
)

// packWriter puts results in job order (take) into the stream s, appends
// their report lines to rep (buffered in w), checkpoints every `every`
// documents with prog, and checks that every document carries the
// TRANSFORMED-BY records tb of the first one.
type packWriter struct {
	s       *pack.Stream
	rep     *os.File
	w       *bufio.Writer
	prog    packProgress
	every   int64
	since   int64
	tb      string
	next    int64
	pending map[int64]packResult
}
