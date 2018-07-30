package utli

import (
	"bytes"
	"unicode"
)

func regexCaseIns(fanhao string) string {
	var buffer bytes.Buffer

	for _, c := range fanhao {
		if isAlpha(c) {
			buffer.WriteRune('[')
			buffer.WriteRune(unicode.ToLower(c))
			buffer.WriteRune(unicode.ToUpper(c))
			buffer.WriteRune(']')
		} else {
			buffer.WriteRune(c)
		}
	}

	return buffer.String()
}

func isAlpha(c rune) bool {
	return (c >= 'a' || c <= 'z') && (c >= 'A' || c <= 'Z')
}
