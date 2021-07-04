package palindrome

import (
	"unicode/utf8"
)

func IsPalindrome(s string) bool {
	reversed := reverse(s)
	if s != reversed {
		return false
	}
	return true
}

func reverse(s string) string {
	totalLength := len(s)
	buffer := make([]byte, totalLength)
	for i := 0; i < totalLength; {
		r, size := utf8.DecodeRuneInString(s[i:])
		i += size
		utf8.EncodeRune(buffer[totalLength-i:], r)
	}
	return string(buffer)
}
