package main

import (
	"testing"
)

func Test_wateringPlants(t *testing.T) {
	result1 := wateringPlants([]int{2, 2, 3, 3}, 5)
	if result1 != 14 {
		t.Errorf("Expected result to be 14, but got %d", result1)
	}
	result2 := wateringPlants([]int{1, 1, 1, 4, 2, 3}, 4)
	if result2 != 30 {
		t.Errorf("Expected result to be 14, but got %d", result2)
	}
	result3 := wateringPlants([]int{7, 7, 7, 7, 7, 7, 7}, 8)
	if result3 != 49 {
		t.Errorf("Expected result to be 14, but got %d", result3)
	}
}
