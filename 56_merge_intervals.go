package main

import (
	"math"
	"sort"
)

func mergeIntervals(a [][]int) [][]int {
	// sort the array from first index small to large
	sort.Slice(a, func(i, j int) bool { return a[i][0] < a[j][0] })

	// initialize empty array and push first index of a
	ans := [][]int{}
	ans = append(ans, a[0])

	for i := range a {
		// if the last index of ans at index 1 is >= beginning of a at i - there is overlap
		if ans[len(ans)-1][1] >= a[i][0] {
			// merge the intervals
			ans[len(ans)-1][1] = int(math.Max(float64(ans[len(ans)-1][1]), float64(a[i][1])))

		} else {
			// push interval as no overlap
			ans = append(ans, a[i])
		}
	}

	return ans
}
