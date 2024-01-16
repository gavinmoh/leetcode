package main

import "math"

func divide(dividend int, divisor int) int {
	sign := -1
	if (dividend < 0 && divisor < 0) || (dividend > 0 && divisor > 0) {
		sign = 1
	}

	dividend = abs(dividend)
	divisor = abs(divisor)
	quotient := 0

	for n := 31; n >= 0; n-- {
		temp := divisor << n
		if temp <= dividend {
			dividend -= temp
			quotient += 1 << n
		}
	}

	result := quotient * sign
	if result > math.MaxInt32 {
		return math.MaxInt32
	} else if result < math.MinInt32 {
		return math.MinInt32
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
