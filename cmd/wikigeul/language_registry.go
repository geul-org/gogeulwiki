//ff:func feature=cli type=loader control=sequence
//ff:what languageRegistry — the GEUL Language Registry from one version of the LANGUAGE-SHORT-CODE codec spec document, with the version and content hash it came from
package main

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/inspect"
	"github.com/geul-org/gogeul/pkg/resolve"

	"github.com/geul-org/wikigeul/internal/port"
)

// registryDoc is the LANGUAGE-SHORT-CODE codec spec document (GEUL Issuer
// 0x0000, document ID = its TAG WORD 0x0848; codec-spec 규격 문서의 위치).
var registryDoc = []core.Word{0, 0, 0, 0x0848}

// languageRegistry reads version v of the LANGUAGE-SHORT-CODE codec spec
// document from dir (v 0: the newest neither deprecated nor deleted, fixed
// once here for the whole run) through the GEUL library's resolve path, and
// returns the registry its ENUMs and BINDs declare
// (codecspec.Spec.Languages, registryOf) and the version it read as
// `0x0000/0x0000000000000848@n 0x<content hash>`, which the run prints and
// the report records. A document without registry is an error.
func languageRegistry(dir resolve.Dir, v uint32) (port.Languages, string, error) {
	h, err := dir.Log([]core.Word{0}, registryDoc)
	if err != nil {
		return nil, "", fmt.Errorf("language registry: %w", err)
	}
	e, ok := registryVersion(h, v)
	if !ok {
		return nil, "", fmt.Errorf("language registry: no version %d of 0x0000/%s", v, core.FormatID(registryDoc))
	}
	b, err := dir.At([]core.Word{0}, registryDoc, e.Version)
	if err != nil {
		return nil, "", fmt.Errorf("language registry: %w", err)
	}
	doc, _ := inspect.Decode(b) // At checked the content hash; a document that does not decode declares no code
	langs, err := registryOf(codecspec.Collect(bodyRecords(doc)).Languages())
	if err != nil {
		return nil, "", err
	}
	return langs, fmt.Sprintf("0x0000/%s@%d 0x%x", core.FormatID(registryDoc), e.Version, e.Hash), nil
}
