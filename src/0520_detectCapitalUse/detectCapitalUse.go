package main

import "fmt"

func main() {
	fmt.Println(detectCapitalUse("Laaa"))
}

func detectCapitalUse(word string) bool {
	count, lenW := 0, len(word)
	for index, v := range word {
		if v <= 'Z' && v >= 'A' {
			if count < index {
				return false
			}
			count++
		}
	}
	if count == lenW || count <= 1 {
		return true
	}
	return false
}
