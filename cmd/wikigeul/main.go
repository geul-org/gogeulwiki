//ff:func feature=cli type=command control=sequence
//ff:what main — wikigeul entrypoint: run, print the error and exit 1 on failure, 3 when a strict run stopped at an anomaly
package main

import (
	"errors"
	"fmt"
	"os"
)

// main runs wikigeul (README) and exits 1 on an error, 3 when a strict run
// stopped at an anomaly (anomalyError).
func main() {
	err := run(os.Args[1:])
	if err == nil {
		return
	}
	fmt.Fprintln(os.Stderr, err)
	if errors.As(err, &anomalyError{}) {
		os.Exit(3)
	}
	os.Exit(1)
}
