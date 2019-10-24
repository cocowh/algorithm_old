package main

import "fmt"

func main() {
	fmt.Println(convertToTitleByChar(27))
}

func convertToTitle(n int) string {
	m := map[int]byte{
		1:  'A',
		2:  'B',
		3:  'C',
		4:  'D',
		5:  'E',
		6:  'F',
		7:  'G',
		8:  'H',
		9:  'I',
		10: 'J',
		11: 'K',
		12: 'L',
		13: 'M',
		14: 'N',
		15: 'O',
		16: 'P',
		17: 'Q',
		18: 'R',
		19: 'S',
		20: 'T',
		21: 'U',
		22: 'V',
		23: 'W',
		24: 'X',
		25: 'Y',
		26: 'Z',
	}
	ret := []byte{}
	for n > 0 {
		if n%26 == 0 {
			ret = append([]byte{'Z'}, ret...)
			n = n/26 - 1
		} else {
			ret = append([]byte{m[n%26]}, ret...)
			n /= 26
		}
		fmt.Println(string(ret), n)
	}
	return string(ret)
}

func convertToTitleByString(n int) string {
	m := "#ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ret := []byte{}
	for n > 0 {
		fmt.Println(n % 26)
		if n%26 == 0 {
			ret = append([]byte{'Z'}, ret...)
			n = n/26 - 1
		} else {
			ret = append([]byte{m[n%26]}, ret...)
			n /= 26
		}
		fmt.Println(string(ret), n)
	}
	return string(ret)
}

func convertToTitleByChar(n int) string {
	ret := []byte{}
	for n > 0 {
		fmt.Println(n % 26)
		if n%26 == 0 {
			ret = append([]byte{'Z'}, ret...)
			n = n/26 - 1
		} else {
			ret = append([]byte{byte(n%26 + 'A' - 1)}, ret...)
			n /= 26
		}
		fmt.Println(string(ret), n)
	}
	return string(ret)
}
