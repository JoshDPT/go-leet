package main

import "testing"

func containsDuplicate_test(t *testing.T) {
	result1 := containsDuplicate([]int{1, 2, 3, 4, 4})
	if result1 != true {
		t.Errorf("Expected result to be true, but got %t", result1)
	}

	result2 := containsDuplicate([]int{1, 2, 3, 4, 5})
	if result2 != false {
		t.Errorf("Expected result to be false, but got %t", result2)
	}
}
