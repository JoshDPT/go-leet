package main

func intSum(ints []int) int {
	sum := 0
	for i := range ints {
		sum += ints[i]
	}
	return sum
}
