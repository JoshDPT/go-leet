package main

import (
	"testing"
)

func Test_commonPrefix (t *testing.T) {
	arr := []string{
		"dog",
		"dong",
		"doom",
		"don",
	}
	result := commonPrefix(arr)
	if result != "do" {
		t.Errorf("Expected result to be do, but got %s", result)
	}
	arr2 := []string{
		"flowbite",
		"flow",
		"flower",
	}
	result2 := commonPrefix(arr2)
	if result2 != "flow" {
		t.Errorf("Expected result to be do, but got %s", result2)
	}
}