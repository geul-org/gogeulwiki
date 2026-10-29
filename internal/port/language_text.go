//ff:func feature=port type=encoder control=sequence
//ff:what languageText — a UTF8-LANGUAGE-TEXT Value from a language code (Wikimedia codes mapped) and text, its language a LANGUAGE-SHORT-CODE when the registry numbers the tag; false when the code is not a BCP 47 tag in form
package port

import (
	"regexp"

	"github.com/geul-org/gogeul/pkg/core"
)

var bcp47Form = regexp.MustCompile(`^[A-Za-z]{2,8}(-[A-Za-z0-9]{1,8})*$`)

// languageText builds the UTF8-LANGUAGE-TEXT of text in language lang. A
// Wikimedia code that is not BCP 47 is first mapped to its tag
// (wikimediaLanguage). The tag must then have the form of a BCP 47 tag
// (letters, then hyphen-separated subtags); only the form is checked. The
// language record is the tag's LANGUAGE-SHORT-CODE when langs (the GEUL
// Language Registry) numbers it, else its LANGUAGE-BCP47-CODE
// (languageRecord).
func languageText(langs Languages, lang, text string) (core.Record, bool) {
	lang = wikimediaLanguage(lang)
	if !bcp47Form.MatchString(lang) {
		return core.Record{}, false
	}
	return core.EncodeLanguageText(languageRecord(langs, lang), []byte(text)), true
}
