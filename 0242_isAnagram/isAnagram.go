package _242_isAnagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	if len(s) == 0 {
		return true
	}
	m := make([]int, 26)
	for index, v := range s {
		m[v - 'a']++
		m[t[index] - 'a']--
	}
	for _,v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
