package main

import (
	"sort"
)

func lastStone(rocks []int) int {
	for len(rocks) > 1 {
		sort.Ints(rocks)
		big := rocks[len(rocks)-1]
		rocks = rocks[:len(rocks)-1]
		big = big - rocks[len(rocks)-1]
		if big > 0 {
			rocks[len(rocks)-1] = big
		} else {
			rocks = rocks[:len(rocks)-1]
		}
	}
	if len(rocks) > 0 {
		return rocks[0]
	} else {
		return 0
	}
}
