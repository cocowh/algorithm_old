package _005_longestPalindrome

func longestPalindrome(s string) string {
	length := len(s)

	if length < 2 {
		return s
	}

	start, maxLen := 0, 1
	for i := 0; i < length; {
		if length-i <= maxLen/2 {
			break
		}
		b, e := i, i
		for e < length-1 && s[e+1] == s[e] {
			e++
		}
		i = e + 1
		for e < length-1 && b > 0 && s[e+1] == s[b-1] {
			e++
			b--
		}
		newLen := e + 1 - b
		if newLen > maxLen {
			start = b
			maxLen = newLen
		}
	}
	return s[start : start+maxLen]
}
