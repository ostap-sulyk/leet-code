package mergestringsalternately

func MergeAlternately(word1, word2 string) string {
	l1, l2 := len(word1), len(word2)
	size := l1 + l2
	result := make([]byte, size)
	i, j, k := 0, 0, 0

	for i < l1 || j < l2 {
		if i < l1 {
			result[k] = word1[i]
			i++
			k++
		}
		if j < l2 {
			result[k] = word2[j]
			j++
			k++
		}
	}

	return string(result)
}
