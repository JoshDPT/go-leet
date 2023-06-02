package main

import (
	"testing"
)

func Test_isNumInList (t *testing.T) {
	result := isNumInList([]int{1,2,3,4,5}, 5)
	if result != true {
		t.Errorf("Expected result to be true, but got %t", result)
	}

	result1 := isNumInList([]int{3,3,3,3}, 5)
	if result1 != false {
		t.Errorf("Expected result1 to be false, but got %t", result1)
	}

	result2 := isNumInList([]int{4,33,-10,8}, -10)
	if result2 != true {
		t.Errorf("Expected result2 to be true, but got %t", result2)
	}

	result3 := isNumInList(nil, 5)
	if result3 != false {
		t.Errorf("Expected result3 to be false, but got %t", result3)
	}

	result4 := isNumInList([]int{}, 5)
	if result4 != false {
		t.Errorf("Expected result4 to be false, but got %t", result4)
	}
}