package _006_convert

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	length := len(s)
	gap := 2*numRows - 2
	var ret string
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < length; j += gap {
			ret += string(s[i+j])
			if i != 0 && i != numRows-1 && j+gap-i < length {
				ret += string(s[j+gap-i])
			}
		}
	}
	return ret
}
