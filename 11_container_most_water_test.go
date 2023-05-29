package main

import (
	"testing"
)

func Test_containerMostWater(t *testing.T) {
	result := containerMostWater([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	if result != 49 {
		t.Errorf("Expected result to be 49, but got %v", result)
	}

}
