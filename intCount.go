package main

func intCount(arr []int) map[int]int {
	countMap := make(map[int]int)
	for _, num := range arr {
		countMap[num]++
	}
	return countMap
}
