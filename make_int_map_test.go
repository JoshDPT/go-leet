package main

import (
	"testing"
	"reflect"
)

func Test_makeMap(t *testing.T) {
	arr := []int{1, 2, 3, 3, 3}
	expected := map[int]int{
		1:1,
		2:2,
		3:3,
	}
	result := makeMap(arr)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}
}