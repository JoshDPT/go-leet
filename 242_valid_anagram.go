package main

func isAnagram(s string, t string) bool {
	one := countBy(s)
	two := countBy(t)
	for k := range one {
		if one[k] != two[k] {
			return false
		}
	}
	for k := range two {
		if one[k] != two[k] {
			return false
		}
	}
	return true
}

func countBy(s string) map[rune]int {
	m := make(map[rune]int)
	for _, v := range s {
		if val, ok := m[v]; ok {
			m[v] = val + 1
		} else {
			m[v] = 1
		}
	}
	return m
}
