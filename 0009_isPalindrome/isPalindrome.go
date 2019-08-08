package _009_isPalindrome

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if 0 <= x && x < 9 {
		return true
	}
	reverse := 0
	temp := x
	for temp != 0 {
		rem := temp % 10 //求余
		temp /= 10
		reverse = reverse*10 + rem
	}
	if reverse ^ x == 0 {
		return true
	} else {
		return false
	}
}