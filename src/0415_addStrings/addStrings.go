package main

import "fmt"

func main() {
	fmt.Println(addStrings("1", "9"))
}

func addStrings(num1 string, num2 string) string {
	res, gap, i, j, carry := []byte{}, int('0'-0), len(num1)-1, len(num2)-1, 0
	for i >= 0 || j >= 0 || carry > 0 {
		n1, n2 := 0, 0
		if i >= 0 {
			n1 = int(num1[i]) - gap
		}
		if j >= 0 {
			n2 = int(num2[j]) - gap
		}
		temp := n1 + n2 + carry
		carry = temp / 10
		res = append([]byte{byte(temp % 10 + gap)},res...)
		i--
		j--
	}
	return string(res)
}
