package main

import "strconv"

func digitSum(n int) int {
	numStr := strconv.Itoa(n)
	sum := 0

	for _, digit := range numStr {
		digitInt, _ := strconv.Atoi(string(digit))
		sum = sum + digitInt
	}

	return sum
}
