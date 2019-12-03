package _010_isMatch


func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}
	firstMatch := s != "" && (s[0] == p[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:])  || (firstMatch && isMatch(s[1:], p))
	} else {
		return firstMatch && isMatch(s[1:], p[1:])
	}
}

