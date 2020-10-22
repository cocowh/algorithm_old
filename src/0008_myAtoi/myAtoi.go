package _008_myAtoi

import (
	"math"
	"strings"
)

func myAtoi(str string) int {

	str = strings.Trim(str, " ")
	len := len(str)
	if len == 0 || str == "" {
		return 0
	}
	i := 0
	result := 0
	flag := true
	if str[i] == '+' {
		i++
	} else if str[i] == '-' {
		i++
		flag = false
	} else if str[i] < '0' || str[i] > '9' {
		return 0
	}
	count := 0
	for ; i < len; i++ {
		if str[i] >= '0' && str[i] <= '9' {
			result = result*10 + int(str[i]) - int('0')
			if result != 0 {
				count++
			}
			if count > 10 {
				if flag {
					return math.MaxInt32
				} else {
					return math.MinInt32
				}
			}
		} else {
			break
		}
	}
	if !flag {
		result = -result
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		if flag {
			result = math.MaxInt32
		} else {
			result = math.MinInt32
		}
	}
	return result
}
