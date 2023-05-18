package main

import "testing"

func Test_subtractProductAndSum(t *testing.T) {
	result := subtractProductAndSum(1234)
	if result != -14 {
		t.Errorf("Expected result to be -14, but got %d", result)
	}
}
