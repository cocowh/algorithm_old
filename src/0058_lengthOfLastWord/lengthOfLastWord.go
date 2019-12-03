package _058_lengthOfLastWord

func lengthOfLastWord(s string) int {
	length := len(s)
	count,flag := 0,false
	for i := length - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if flag {
				return count
			} else {
				continue
			}
		} else if ('a' <= s[i] && s[i] <= 'z') || ('A' <= s[i] && s[i] <= 'Z') {
			count ++
			flag = true
		} else {
			count = 0
			flag =false
		}
	}
	if flag {
		return  count
	} else {
		return 0
	}
}