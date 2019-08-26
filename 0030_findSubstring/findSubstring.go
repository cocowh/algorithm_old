package _030_findSubstring

func findSubstring(s string, words []string) []int {
	ret,wordNum := []int{},len(words)
	if wordNum == 0 {
		return ret
	}
	wordLen := len(words[0])
	allWards := make(map[string]int)
	for _,word := range words {
		if _,ok := allWards[word];ok {
			allWards[word]++
		} else {
			allWards[word] = 0
		}
	}
	for j := 0;j < wordLen; j++ {
		num := 0
		hasWords := make(map[string]int)
		for i := j; i < len(s) - wordNum * wordLen + 1;i = i + wordLen {
			hasRemoved := false
			for num < wordNum {
				word := s[i + num * wordLen : i + (num + 1) * wordLen]
				if _, ok := allWards[word];ok{
					if _, ok := hasWords[word]; ok {
						hasWords[word]++
					} else {
						hasWords[word] = 0
					}
					if hasWords[word] > allWards[word] {
						hasRemoved = true
						removeNum := 0
						for hasWords[word] > allWards[word] {
							firstWord := s[i + removeNum * wordLen : i + (removeNum + 1) * wordLen]
							hasWords[firstWord]--
							removeNum++
						}
						num = num - removeNum + 1
						i = i + (removeNum - 1) * wordLen
						break
					}
				} else {
					hasWords = make(map[string]int)
					i = i + num * wordLen
					num = 0
					break
				}
				num++
			}
			if num == wordNum {
				ret = append(ret,i)
			}
			if num > 0 && !hasRemoved {
				firstWord := s[i : i + wordLen]
				hasWords[firstWord]--
				num--
			}
		}
	}
	return ret
}