package main

func findLucky(arr []int) int {
	countMap := intCount(arr)
	maxLucky := -1
	for num, count := range countMap {
		if num == count && num > maxLucky {
			maxLucky = num
		}
	}
	return maxLucky
}
