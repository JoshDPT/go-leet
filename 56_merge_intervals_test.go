package main

import (
	"reflect"
	"testing"
)

func Test_mergeIntervals(t *testing.T) {
	result0 := mergeIntervals([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})

	expected0 := [][]int{{1, 6}, {8, 10}, {15, 18}}

	if !reflect.DeepEqual(result0, expected0) {
		t.Errorf("Expected result to be %v, but got %v", expected0, result0)
	}

	result1 := mergeIntervals([][]int{{1, 4}, {4, 5}})
	expected1 := [][]int{{1, 5}}

	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("expected result to be %v, but got %v", expected1, result1)
	}

	result2 := mergeIntervals([][]int{{1, 4}, {0, 4}})
	expected2 := [][]int{{0, 4}}

	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("expected result to be %v, but got %v", expected2, result2)
	}

}
