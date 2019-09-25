package main

import "fmt"

func main() {
	instance := ConstructorNative(4)
	fmt.Println(instance.Get(1))
	instance.Put(520,1314)
	fmt.Println(instance.Get(520))
	instance.Put("string","I am a string")
	fmt.Println(instance.Get("string"))
}
