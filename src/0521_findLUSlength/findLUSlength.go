package main

import "fmt"

func main() {
	fmt.Println(findLUSlength("a","ab"))
}

func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	if len(a) > len(b) {
		return len(a)
	}
	return len(b)
}