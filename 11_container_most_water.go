package main

import (
	"math"
)

func containerMostWater (a []int) int{
	left, right := 0, len(a)-1
	max := 0.0

	for left < right {
		max = math.Max(max, (math.Min(float64(a[left]), float64(a[right])) * (float64(right - left))))

		if a[left] < a[right] {
			left++
		} else {
			right--
		}
	}
	
	return int(max)
}