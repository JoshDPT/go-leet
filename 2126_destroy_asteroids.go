package main

import (
	"sort"
)

func destroyAsteroids(planet int, asteroids []int) bool {

	sort.Ints(asteroids)

	for _, asteroid := range asteroids {
		if planet >= asteroid {
			planet += asteroid
		} else {
			return false
		}
	}

	return true
}
