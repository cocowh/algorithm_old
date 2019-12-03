package _028_strStr

func strStr(haystack string, needle string) int {
	needleLen,haystackLen := len(needle),len(haystack)
	if needleLen == 0 {
		return 0
	}
	if needleLen > haystackLen {
		return -1
	}
	for i := 0; i < haystackLen; i++ {
		if i + needleLen <= haystackLen && haystack[i:i + needleLen] == needle{
			return i
		} else if i + needleLen > haystackLen {
			return -1
		}
	}
	return -1
}