package main

import (
	"testing"
	"reflect"
)

func Test_rotateImage(t *testing.T) {
	m := [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	expected := [][]int{
		{7,4,1},
		{8,5,2},
		{9,6,3},
	}
	result := rotateImage(m)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result to be %v, but got %v", expected, result)
	}
}