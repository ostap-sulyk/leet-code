package countvowelstringsinranges

import (
	"slices"
)

func vowelStrings(words []string, queries [][]int) []int {
	vowels := []byte{'a', 'e', 'i', 'o', 'u'}
	prefix_cnt := []int{0}
	res := make([]int, len(queries))
	curr := 0

	for _, word := range words {
		if slices.Contains(vowels, word[0]) && slices.Contains(vowels, word[len(word)-1]) {
			curr += 1
		}
		prefix_cnt = append(prefix_cnt, curr)
	}

	for i, query := range queries {
		res[i] = prefix_cnt[query[1]+1] - prefix_cnt[query[0]]
	}
	return res
}
