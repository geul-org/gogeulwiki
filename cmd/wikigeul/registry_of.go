//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what registryOf — port.Languages from a codec spec's language ENUMs: lowercase tag to 1-WORD code
package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/geul-org/gogeul/pkg/core"

	"github.com/geul-org/wikigeul/internal/port"
)

// registryOf turns enums (code in canonical hex `0xNNNN` to BCP 47 tag,
// codecspec.Spec.Languages) into port.Languages, keyed by the tag in lower
// case. An empty registry, a code that is not 1 WORD and a tag named twice
// (case ignored) are errors.
func registryOf(enums map[string]string) (port.Languages, error) {
	if len(enums) == 0 {
		return nil, errors.New("language registry: the LANGUAGE-SHORT-CODE spec declares no code")
	}
	out := port.Languages{}
	for code, tag := range enums {
		n, err := strconv.ParseUint(strings.TrimPrefix(code, "0x"), 16, 16)
		key := strings.ToLower(tag)
		_, dup := out[key]
		if err != nil || !strings.HasPrefix(code, "0x") || dup {
			return nil, fmt.Errorf("language registry: code %s tag %q", code, tag)
		}
		out[key] = core.Word(n)
	}
	return out, nil
}
