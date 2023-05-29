package main

import (
	"testing"
)

func Test_destroy_asteroids(t *testing.T) {
	result := destroyAsteroids(10, []int{3, 9, 19, 5, 21})
	if result != true {
		t.Errorf("Expected result to be true, but got %t", result)
	}
	result2 := destroyAsteroids(5, []int{4, 9, 23, 4})
	if result2 != false {
		t.Errorf("Expected result2 to be false, but got %t", result2)
	}
	result3 := destroyAsteroids(10, []int{4, 9, 23, 600})
	if result3 != false {
		t.Errorf("Expected result3 to be false, but got %t", result3)
	}

}
