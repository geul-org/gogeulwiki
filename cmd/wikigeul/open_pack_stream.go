//ff:func feature=cli type=command control=sequence
//ff:what openPackStream — the Issuer's pack stream: resumed from its checkpoint (same run only) or new with a dictionary trained on samples
package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/geul-org/gogeul/pkg/pack"

	"github.com/geul-org/wikigeul/internal/port"
)

// openPackStream opens the stream for set. With -resume it continues the
// checkpointed stream; the checkpoint must come from the same run (sameRun:
// the same commits in TRANSFORMED-BY, registry, dump, range or list,
// RECORDED-AT) — a resumed run that differed would mix documents. Otherwise
// it starts a new set: zstd with Workers concurrent encoders and a
// dictionary trained on dictSamples, created at RECORDED-AT.
func openPackStream(set string, c config, prov port.Provenance, db *sql.DB, ids []packJob, idsHash string) (*pack.Stream, packProgress, error) {
	want := progressOf(c, idsHash)
	o := pack.StreamOptions{Build: pack.BuildOptions{Codec: &pack.Zstd{Workers: c.Workers}}, Issuers: []string{"0x00010000"}, Created: c.RecordedAt}
	if c.Resume {
		s, err := pack.ResumeStream(set, o)
		if err != nil {
			return nil, want, err
		}
		var p packProgress
		err = json.Unmarshal(s.Meta(), &p)
		if c.RecordedAt == 0 {
			want.RecordedAt = p.RecordedAt // -recorded-at left out: the checkpoint's
		}
		if err != nil || !sameRun(p, want) {
			return nil, want, errors.New("wikigeul: -resume: the checkpoint is from another run (commits, registry, dump, range, ID list or -recorded-at differ)")
		}
		return s, p, nil
	}
	samples, err := dictSamples(db, c, prov, ids)
	if err != nil {
		return nil, want, fmt.Errorf("dictionary samples: %w", err)
	}
	o.Samples = samples
	s, err := pack.NewStream(set, o)
	return s, want, err
}
