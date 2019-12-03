package main

import "fmt"

func main() {
	fmt.Printf("%b\n",111)
	fmt.Println(hammingWeight(111))
}

func hammingWeight(num uint32) int {
	count := 0
	for i := 0; i < 32; i++ {
		if num & 1 == 1 {
			count++
		}
		num>>=1
	}
	return count
}
