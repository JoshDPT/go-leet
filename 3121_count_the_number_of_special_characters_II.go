package main

import (
	"unicode"
)

// Runtime: 87 ms, faster than 41.99% of Go online submissions for Count the Number of Special Characters II.
// Memory Usage: 6.9 MB, less than 85.71% of Go online submissions for Count the Number of Special Characters II.

func numberOfSpecialCharsTwo(word string) int {
	lowerSet := make(map[rune]int)
	upperSet := make(map[rune]int)

	for i, r := range word {
		if unicode.IsLetter(r) && unicode.IsLower(r) {
			// if its a letter and lowercase - put the most recent index
			lowerSet[unicode.ToUpper(r)] = i
		} else if unicode.IsLetter(r) {
			// if it is upper - only put and index if it is not in the set
			_, ok := upperSet[r]
			if !ok {
				upperSet[r] = i
			}
		}
	}

	var count int
	for r := range upperSet {
		if _, ok := upperSet[r]; ok {
			if _, ok := lowerSet[r]; ok {
				if lowerSet[r] < upperSet[r] {
					count += 1
				}
			}
		}
	}
	return count
}
