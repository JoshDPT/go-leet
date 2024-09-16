package main

import (
	"strings"
	"unicode"
)

func validPali(s string) bool {
	lower := convStrLowerAlphaNum(s)
	left := 0
	right := len(lower) - 1

	for left < right {
		if lower[left] != lower[right] {
			return false
		}
		left += 1
		right -= 1
	}
	return true
}

func convStrLowerAlphaNum(s string) string {
	str := ""
	lower := strings.ToLower(s)
	for _, right := range lower {
		if unicode.IsLetter(right) || unicode.IsDigit(right) {
			str += string(right)
		}
	}
	return str
}
