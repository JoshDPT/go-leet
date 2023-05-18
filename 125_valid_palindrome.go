package main

import (
	"strings"
	"unicode"
)

func validPali(s string) bool {
	lower := convStrLowerAlphaNum(s)
	l := 0
	r := len(lower) - 1

	for l < r {
		if lower[l] != lower[r] {
			return false
		}
		l += 1
		r -= 1
	}
	return true
}

func convStrLowerAlphaNum(s string) string {
	str := ""
	lower := strings.ToLower(s)
	for _, r := range lower {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			str += string(r)
		}
	}
	return str
}
