package main

import (
	"fmt"
	"strconv"
)

func main()  {
	fmt.Println(countAndSayBetter(3))
}
func countAndSay(n int) string {
	res := "1"
	for i := 1;i < n; i++ {
		resTemp := ""
		for j:= 0;j < len(res); {
			count := 1
			for k := j + 1;k < len(res) && res[k] == res[j];k++{
				count++
			}
			resTemp = resTemp + strconv.Itoa(count) + string(res[j])
			j += count
		}
		res = resTemp
	}
	return res
}

func countAndSayBetter(n int) string {
	res := []byte{'1'}
	for i := 1;i < n; i++ {
		resTemp := []byte{}
		for j:= 0;j < len(res); {
			count := 1
			for k := j + 1;k < len(res) && res[k] == res[j];k++{
				count++
			}
			resTemp = append(resTemp, byte(count + 48), res[j])
			j += count
		}
		res = resTemp
	}
	return string(res)
}