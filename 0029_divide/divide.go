package _029_divide

import "math"

func divide(dividend int, divisor int) int {
	dividendAbs := int(math.Abs(float64(dividend)))
	divisorAbs := int(math.Abs(float64(divisor)))
	quotient := 0
	for i := uint32(31);i >= 0 && dividendAbs >= divisorAbs; i-- {
		if (dividendAbs >> i) >= divisorAbs {
			quotient = quotient + ( 1 << i)
			dividendAbs = dividendAbs - (divisorAbs << i)
		}
	}
	if dividend ^ divisor < 0 {
		quotient = -quotient
	}
	if quotient > math.MaxInt32 || quotient < math.MinInt32 {
		return  math.MaxInt32
	}
	return quotient
}