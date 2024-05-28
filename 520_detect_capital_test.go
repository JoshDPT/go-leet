package main

import (
	"testing"
)

func Test_detectCapitalUse(t *testing.T) {
	result1 := detectCapitalUse("USA")
	if result1 != true {
		t.Errorf("Expected result1 to be true, but got %v", result1)
	}
	result2 := detectCapitalUse("FLaG")
	if result2 != false {
		t.Errorf("Expected result2 to be false, but got %v", result2)
	}
}
