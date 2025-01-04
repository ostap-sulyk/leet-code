package validnumber

func IsNumber1(s string) bool {
	seenDigit, seenExponent, seenDot := false, false, false
	var curr byte
	for i := 0; i < len(s); i++ {
		curr = s[i]
		if '0' <= curr && curr <= '9' {
			seenDigit = true
		} else if curr == '+' || curr == '-' {
			if i > 0 && s[i-1] != 'e' && s[i-1] != 'E' {
				return false
			}
		} else if curr == 'e' || curr == 'E' {
			if seenExponent || !seenDigit {
				return false
			}
			seenExponent = true
			seenDigit = false
		} else if curr == '.' {
			if seenDot || seenExponent {
				return false
			}
			seenDot = true
		} else {
			return false
		}
	}
	return seenDigit
}

// aproach two
func IsNumber2(s string) bool {
	return false
}