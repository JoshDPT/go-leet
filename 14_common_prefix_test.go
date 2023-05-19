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
}