package maximumscoreaftersplittingastring

func maxScore(s string) int {
	count0, count1, res := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			count1++
		}
	}

	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			count0++
		} else {
			count1--
		}
		if count0+count1 > res {
			res = count0 + count1
		}
	}
	return res
}
