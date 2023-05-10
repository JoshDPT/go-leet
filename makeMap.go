package main

func makeMap(nums []int) map[int]int {
	numsMap := make(map[int]int)
	for _, n := range nums {
		numsMap[n] = n
	}
	return numsMap
}
