package main

import "unicode"

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Detect Capital.
// Memory Usage: 2.3 MB, less than 68.42% of Go online submissions for Detect Capital.

func detectCapitalUse(word string) bool {
	upper := IsUpper(word)
	lower := IsLower(word)
	return upper || lower
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	for i, r := range s {
		if i != 0 {
			if !unicode.IsLower(r) && unicode.IsLetter(r) {
				return false
			}
		}
	}
	return true
}
