package mergestringsalternately

func MergeAlternately(word1, word2 string) string {
	l := len(word1) + len(word2)
	result := make([]byte, l)
	i, j, k := 0, 0, 0

	for i < len(word1) || j < len(word2) {
		if i < len(word1) {
			result[k] = word1[i]
			i++
			k++
		}
		if j < len(word2) {
			result[k] = word2[j]
			j++
			k++
		}
	}

	return string(result)
}
