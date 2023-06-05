package main

import (
	"strconv"
)

func divisibilityArray(word string, m int) []int {
	n := len(word)
	div := make([]int, n)
	current := 0

	for i := 0; i < n; i++ {
		digit, _ := strconv.Atoi(string(word[i]))
		current = (current*10 + digit) % m
		if current == 0 {
			div[i] = 1
		}
	}

	return div
}
