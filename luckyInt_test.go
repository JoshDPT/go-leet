package main

import "testing"

func TestCalc1(t *testing.T) {
	result := findLucky([]int{1, 2, 3, 4, 2, 5})
	if result != 2 {
		t.Errorf("Expected result to be 2, but got %d", result)
	}
}
