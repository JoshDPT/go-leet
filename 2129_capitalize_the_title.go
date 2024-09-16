package main

import (
	"strings"
)

func capitalizeTitle(title string) string {
	allLower := strings.ToLower(title)
	words := strings.Split(allLower, " ")
	for i, word := range words {
		if len(word) > 2 {
			titleCase := strings.ToUpper(word[:1]) + word[1:]
			words[i] = titleCase
		}
	}
	return strings.Join(words, " ")
}
