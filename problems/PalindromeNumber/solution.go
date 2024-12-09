package palindromenumber

func isPalindrome(x int) bool {
	if x <= 0 || x%10 == 0 {
		return false
	}

	original, reversed := x, 0

	for x > 0 {
		reversed = reversed*10 + x%10
		x = x / 10
	}

	return original == reversed
}
