package main

import "strconv"

func digitProduct(n int) int {
	numStr := strconv.Itoa(n)
	product := 1

	for _, digit := range numStr {
		digitInt, _ := strconv.Atoi(string(digit))
		product = product * digitInt
	}

	return product
}
