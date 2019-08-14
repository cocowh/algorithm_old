package _017_letterCombinations

func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}
	ret := []string{""}
	m := map[byte][]string{
		'2':[]string{"a","b","c"},
		'3':[]string{"d","e","f"},
		'4':[]string{"g","h","i"},
		'5':[]string{"j","k","l"},
		'6':[]string{"m","n","o"},
		'7':[]string{"p","q","r","s"},
		'8':[]string{"t","u","v"},
		'9':[]string{"w","x","y","z"},
	}
	for _,v := range digits{
		copy := []string{}
		retLen := len(ret)
		for j := 0; j < retLen; j++ {
			arrLen := len(m[byte(v)])
			for i := 0; i < arrLen; i++ {
				copy = append(copy, ret[j] + m[byte(v)][i])
			}
		}
		ret = copy
	}
	return ret
}