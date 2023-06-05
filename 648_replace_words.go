package main

import (
	"sort"
	"strings"
)

func replaceWords(dic []string, s string) string {
	words := strings.Fields(s)

	// Define the custom comparison function
	compareFn := func(i, j int) bool {
		return len(dic[i]) < len(dic[j])
	}

	// Sort the slice using the custom comparison function
	sort.Slice(dic, compareFn)

	for i, word := range words {
		for _, pre := range dic {
			if strings.HasPrefix(word, pre) && len(pre) < len(words[i]) {
				words[i] = pre
			}
		}
	}
	return strings.Join(words, " ")
}
