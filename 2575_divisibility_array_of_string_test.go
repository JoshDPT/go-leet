package main

// Input: word = "998244353", m = 3
// Output: [1,1,0,0,0,1,1,0,0]

// Input: word = "1010", m = 10
// Output: [0,1,0,1]

import (
	"reflect"
	"testing"
)

func Test_divisibilityArray(t *testing.T) {
	expected := []int{1, 1, 0, 0, 0, 1, 1, 0, 0}
	result := divisibilityArray("998244353", 3)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}

	expected2 := []int{0, 1, 0, 1}
	result2 := divisibilityArray("1010", 10)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Expected result2 to be %v, but got %v", expected2, result2)
	}
}
