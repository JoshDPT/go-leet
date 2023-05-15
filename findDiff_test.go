package main

import (
	"reflect"
	"testing"
)

func Test_findDifference(t *testing.T) {
	expected := [][]int{
		{1, 3},
		{4, 6},
	}
	result := findDifference([]int{1, 2, 3}, []int{2, 4, 6})
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}
}
