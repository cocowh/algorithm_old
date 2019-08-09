package _012_intToRoman

func intToRoman(num int) string {
	signMap := [4][]string{
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		{"", "M", "MM", "MMM"},
	}
	return signMap[3][num/1000] + signMap[2][num/100%10] + signMap[1][num/10%10] + signMap[0][num%10]
}