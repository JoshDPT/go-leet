package main

import (
	"unicode"
)

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Count the Number of Special Characters I.
// Memory Usage: 2.4 MB, less than 65.98% of Go online submissions for Count the Number of Special Characters I.

func numberOfSpecialChars(word string) int {
	lowerSet := make(map[rune]bool)
	upperSet := make(map[rune]bool)

	for _, r := range word {
		if unicode.IsLetter(r) && unicode.IsLower(r) {
			lowerSet[unicode.ToUpper(r)] = true
		} else if unicode.IsLetter(r) {
			upperSet[r] = true
		}
	}

	var count int

	for r := range upperSet {
		if _, ok := upperSet[r]; ok {
			if _, ok := lowerSet[r]; ok {
				count += 1

			}
		}
	}

	return count
}
