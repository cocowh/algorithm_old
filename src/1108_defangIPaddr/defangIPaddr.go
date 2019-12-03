package main

import "fmt"

func main()  {
	fmt.Println(defangIPaddr("255.100.50.0"))
}

func defangIPaddr(address string) string {
	ret := []byte{}
	for i := 0; i < len(address); i++ {
		if address[i] == '.' {
			ret = append(ret, '[','.',']')
		} else {
			ret = append(ret, address[i])
		}
	}
	return string(ret)
}
