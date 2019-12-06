package main

import "fmt"

func main() {
	fmt.Println(findTheDifference("abcd", "abcde"))
}

func findTheDifference(s string, t string) byte {
	ret := 0
	for _,v := range s {
		ret ^= int(v)
	}
	for _,v := range t {
		ret ^= int(v)
	}
	return byte(ret)
}

func findTheDifferenceByAdd(s string, t string) byte {
	counts,countt := int32(0),int32(0)
	for _,v := range s {
		counts += v
	}
	for _,v := range t {
		countt += v
	}
	return byte(countt-counts)
}