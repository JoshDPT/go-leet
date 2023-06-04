
// Input: dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
// Output: "the cat was rat by the bat"

// Input: dictionary = ["a","b","c"], sentence = "aadsfasf absbs bbab cadsfafs"
// Output: "a a b c"

package main

import "testing"

func Test_replaceWords (t *testing.T) {
	result := replaceWords([]string{"cat","bat","rat"},"the cattle was rattled by the battery")
	if result != "the cat was rat by the bat" {
		t.Errorf("Expected result2 to be (the cat was rat by the bat), but got %v", result)
	}

	result2 := replaceWords([]string{"a", "aa", "aaa", "aaaa"},"a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa")
	if result2 != "a a a a a a a a bbb baba a" {
		t.Errorf("Expected result2 to be (a a a a a a a a bbb baba a), but got %v", result2)
	}
}