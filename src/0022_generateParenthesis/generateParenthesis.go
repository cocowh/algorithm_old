package _022_generateParenthesis

func generateParenthesis(n int) []string {
	ret := []string{}
	getParenthesis(&ret, "", 0, 0, n)
	return ret
}

func getParenthesis(ret *[]string, cur string, left, right, max int)  {
	if len(cur) == max * 2 {
		*ret = append(*ret, cur)
		return
	}
	if left < max {
		getParenthesis(ret,cur + "(", left + 1, right, max)
	}
	if right < left {
		getParenthesis(ret,cur + ")", left, right + 1, max)
	}
}