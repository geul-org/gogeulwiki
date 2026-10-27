//ff:func feature=port type=helper control=selection
//ff:what isTitleByte — whether a byte of a MediaWiki page title stays as is in its page URL (wfUrlencode)
package port

// isTitleByte reports whether c stays unencoded in a MediaWiki page URL, as
// MediaWiki's wfUrlencode leaves it: a letter, digit, - . _ (PHP urlencode)
// or one of ; @ $ ! * ( ) , / ~ : (decoded back by wfUrlencode). Everything
// else, including % & ' + = ? # and every UTF-8 byte, is percent-encoded.
func isTitleByte(c byte) bool {
	switch {
	case c >= 'a' && c <= 'z', c >= 'A' && c <= 'Z', c >= '0' && c <= '9':
		return true
	default:
		return c != 0 && containsByte("-._;@$!*(),/~:", c)
	}
}
