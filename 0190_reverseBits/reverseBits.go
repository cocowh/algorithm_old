package main

import "fmt"

func main() {
	fmt.Printf("%b\n", reverseBitsDiff(111))
}

func reverseBits(num uint32) uint32 {
	ret := uint32(0)
	for i := 0; i < 32; i++ {
		ret <<= 1
		ret += num & 1
		num >>= 1
	}
	return ret
}

func reverseBitsDiff(num uint32) uint32 {
	num = (num >> 16) | (num << 16)
	fmt.Println(num)
	num = ((num & 0xff00ff00) >> 8) | ((num & 0x00ff00ff) << 8)
	fmt.Println(num)
	num = ((num & 0xf0f0f0f0) >> 4) | ((num & 0x0f0f0f0f) << 4)
	fmt.Println(num)
	num = ((num & 0xcccccccc) >> 2) | ((num & 0x33333333) << 2)
	fmt.Println(num)
	num = ((num & 0xaaaaaaaa) >> 1) | ((num & 0x55555555) << 1)
	return num
}