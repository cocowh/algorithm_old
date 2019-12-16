package _657_judgeCircle

func judgeCircle(moves string) bool {
	r, l := 0, 0
	for _, v := range moves {
		switch v {
		case 'R':
			r++
		case 'L':
			r--
		case 'U':
			l++
		case 'D':
			l--
		}
	}
	return r == 0 && l == 0
}
