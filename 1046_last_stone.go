package main

import (
	"sort"
)

func lastStone(stones []int) int {

	for len(stones) > 1 {
		sort.Ints(stones)
		big := stones[len(stones)-1]
		stones = stones[:len(stones)-1]
		big = big - stones[len(stones)-1]
		if big > 0 {
			stones[len(stones)-1] = big
		} else {
			stones = stones[:len(stones)-1]
		}
	}
	if len(stones) > 0 {
		return stones[0]
	} else {
		return 0
	}
}
