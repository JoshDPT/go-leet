package main

import (
	"strings"
	"unicode"
)

func decodeAtIndex(s string, k int) string {
	var final string

	for _, rune := range s {
		if unicode.IsNumber(rune) {
			final = strings.Repeat(final, int(rune)-int('0'))
		} else {
			final += string(rune)
		}
	}
	return string(final[k])
}
