//ff:func feature=cli type=command control=iteration dimension=1
//ff:what writeAnomalies — the anomaly report: comment lines, a header and one TSV line per anomaly (value quoted)
package main

import (
	"os"
	"strings"
)

// anomalyHeader is the column line of the anomaly report.
const anomalyHeader = "id\tkind\tproperty\tvalue\treason\tclass"

// writeAnomalies writes head (comment lines), the header and one line per
// anomaly (anomalyLine) to path, replacing the file.
func writeAnomalies(path, head string, as []anomaly) error {
	var b strings.Builder
	b.WriteString(head + "\n" + anomalyHeader + "\n")
	for _, a := range as {
		b.WriteString(anomalyLine(a) + "\n")
	}
	return os.WriteFile(path, []byte(b.String()), 0o644)
}
