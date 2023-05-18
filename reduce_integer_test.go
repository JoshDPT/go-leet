package main

import "testing"

func Test_reduceInt(t *testing.T) {
	result := reduceInt([]int{1, 2, 3, 4, 5}, func(a, b int) int { c := a + b; return c }, 0)
	if result != 15 {
		t.Errorf("Expected result to be 10, but got %d", result)
	}
}
