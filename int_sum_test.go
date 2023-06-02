package main

import "testing"

func Test_intSum (t *testing.T) {
	result := intSum([]int{1,2,3,4,5})
	if result != 15 {
		t.Errorf("Expected result to be true, but got %v", result)
	}

	result1 := intSum([]int{3,3,3,3})
	if result1 != 12 {
		t.Errorf("Expected result1 to be false, but got %v", result1)
	}

	result2 := intSum([]int{5,50,-10,5})
	if result2 != 50 {
		t.Errorf("Expected result2 to be true, but got %v", result2)
	}

	result3 := intSum(nil)
	if result3 != 0 {
		t.Errorf("Expected result3 to be 0, but got %v", result3)
	}

	result4 := intSum([]int{})
	if result4 != 0 {
		t.Errorf("Expected result4 to be 0, but got %v", result4)
	}
}