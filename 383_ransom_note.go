package main

func ransomNote(note string, magazine string) bool {
	n := countMap(note)
	m := countMap(magazine)

	for key := range n {
		if m[key] < n[key] {
			return false
		}
	}
	return true
}

func countMap(s string) map[rune]int {
	m := make(map[rune]int)
	for _, c := range s {
		if val, ok := m[c]; ok {
			m[c] = val + 1
		} else {
			m[c] = 1
		}
	}
	return m
}
