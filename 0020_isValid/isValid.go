package _020_isValid

func isValid(s string) bool {
	m := map[byte]byte{
		'(':')',
		'{':'}',
		'[':']',
	}
	stack := []byte{}
	for _,v := range s {
		if v == '{' || v == '(' || v == '['  {
			stack = append(stack,byte(v))
		} else if v == '}' && v == ')' || v == ']' {
			if len(stack) > 0 || m[stack[len(stack) - 1]] == byte(v) {
				stack = stack[:len(stack) - 1]
				continue
			} else {
				return false
			}
		}
	}
	return  len(stack) == 0
}