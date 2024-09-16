package main

func longestSubstringNoRepeat(s string) int {
	set := make(map[byte]bool)
	left, max := 0, 0

	for right := 0; right < len(s); right += 1 {
		for set[s[right]] {
			delete(set, s[left])
			left += 1
		}
		set[s[right]] = true

		if len(set) > max {
			max = len(set)
		}
	}

	return max
}
