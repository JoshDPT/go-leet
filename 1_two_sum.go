package main

func twoSum(arr []int, x int) []int {
	m := make(map[int]int)

	for i, v := range arr {
		if val, ok := m[x-v]; ok {
			return []int{val, i}
		}
		m[v] = i
	}

	return []int{0, 0}
}
