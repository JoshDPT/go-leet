package main

import (
	"sort"
	"strings"
)

type Pair struct {
	key rune
	value int
}

func frequencySort (s string) string {
	count := countBy(s)
	pairs := make([]Pair, 0, len(count))

	for key, value := range count {
		pairs = append(pairs, Pair{key, value})
	}

	sort.Slice(pairs, func(i,j int) bool {
		return pairs[i].value > pairs[j].value
	})

	builder := []string{}

	for _,pair := range pairs {
		i := pair.value
		for i > 0 {
			str := string(pair.key)
			builder = append(builder, str)
			i--
		}

	}

	return strings.Join(builder, "")

}