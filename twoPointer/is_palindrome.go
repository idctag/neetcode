package twopointer

import "unicode"

func IsPalindrome(s string) bool {
	l := 0
	r := len(s) - 1

	for l < r {
		if !unicode.IsLetter(rune(s[l])) && !unicode.IsNumber(rune(s[l])) {
			l++
			continue
		}
		if !unicode.IsLetter(rune(s[r])) && !unicode.IsNumber(rune(s[r])) {
			r--
			continue
		}
		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		}
		r--
		l++
	}

	return true
}
