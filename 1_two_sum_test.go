package main

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	expected := []int{1, 2}
	result := twoSum([]int{1, 2, 3}, 5)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}
}
