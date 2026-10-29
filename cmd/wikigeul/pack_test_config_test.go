//ff:func feature=cli type=command control=sequence
//ff:what packTestConfig — a pack-mode config over the fixture schema: small checkpoints, two workers, fixed RECORDED-AT, a temp root
package main

import (
	"strings"
	"testing"
)

// packTestConfig is a pack-mode run of schema's dump 1 into a temp root.
func packTestConfig(t *testing.T, schema string) config {
	a := strings.Repeat("a", 40)
	return config{Root: t.TempDir(), From: 1, To: 100, Schema: schema, Dump: 1, Commit: a, LibCommit: a,
		Out: "pack", Workers: 2, Every: 2, Samples: 2, RecordedAt: 1_800_000_000}
}
