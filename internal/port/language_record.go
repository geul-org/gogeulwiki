//ff:func feature=port type=encoder control=sequence
//ff:what languageRecord — the language Value of a BCP 47 tag: its LANGUAGE-SHORT-CODE when the registry numbers it, else a LANGUAGE-BCP47-CODE in canonical case
package port

import (
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
)

// languageRecord returns the language record of tag (value/text/
// UTF8-LANGUAGE-TEXT): a LANGUAGE-SHORT-CODE when langs has the tag (compared
// without case, as BCP 47 tags are), else a LANGUAGE-BCP47-CODE holding the
// tag in canonical case (canonicalTag). A tag the registry does not number —
// private use, extensions, a tag IANA does not register — is never an error:
// LANGUAGE-BCP47-CODE is its form.
func languageRecord(langs Languages, tag string) core.Record {
	if code, ok := langs[strings.ToLower(tag)]; ok {
		return core.Record{Kind: core.KindValue, Codec: codecLanguageShort, Payload: []core.Word{code}}
	}
	return core.EncodeBytesValue(core.CodecLanguageBCP47Even, []byte(canonicalTag(tag)))
}
