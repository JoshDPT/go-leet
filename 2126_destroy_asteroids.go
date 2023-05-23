package main

import (
	"fmt"
	"sort"
)

func main () {
	arr := []int{3,9,19,5,21}
	fmt.Println(destroyAsteroids(10, arr))
}

func destroyAsteroids(planet int, asteroids []int) bool {

	sort.Ints(asteroids)

	for _,asteroid := range asteroids {
		if planet >= asteroid {
			planet += asteroid
		} else {
			return false
		}
	}

	return true
}