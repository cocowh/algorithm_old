package _014_longestCommonPrefix

func longestCommonPrefix(strs []string) string {
	strArrLen := len(strs)
	if strArrLen == 0 {
		return ""
	}
	if strArrLen == 1 {
		return strs[0]
	}
	minLen := len(strs[0])
	for _,str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}
	for minLen >0 {
		for i := 0; i< strArrLen; i++ {
			if strs[i][:minLen] == strs[0][:minLen] {
				if  i == strArrLen - 1 {
					return strs[0][:minLen]
				} else {
					continue
				}
			} else {
				minLen--
				break
			}
		}
	}
	return  ""
}