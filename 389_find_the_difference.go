package main

import (
	"sort"
)

func findDifference(nums1 []int, nums2 []int) [][]int {
	nums1diff := []int{}
	nums2diff := []int{}
	map1 := makeMap(nums1)
	map2 := makeMap(nums2)
	
	for k := range map1 {
		if _, ok := map2[k]; !ok {
			nums1diff = append(nums1diff, k)
		}
	}

	for k := range map2 {
		if _, ok := map1[k]; !ok {
			nums2diff = append(nums2diff, k)
		}
	}

	sort.Slice(nums1diff, func(i, j int) bool {
		return nums1diff[i] < nums1diff[j]
	})

	sort.Slice(nums2diff, func(i, j int) bool {
		return nums2diff[i] < nums2diff[j]
	})

	diffSlice := [][]int{}

	diffSlice = append(diffSlice, nums1diff)
	diffSlice = append(diffSlice, nums2diff)

	return diffSlice
}
