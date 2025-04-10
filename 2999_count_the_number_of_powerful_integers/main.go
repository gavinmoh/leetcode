package main

import (
	"math"
	"strconv"
)

func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {
	calculate := func(x string) int64 {
		if len(x) < len(s) {
			return 0
		}

		if len(x) == len(s) {
			if x >= s {
				return 1
			}
			return 0
		}

		suffix := x[len(x)-len(s):]
		count := int64(0)
		preLen := len(x) - len(s)

		for i := 0; i < preLen; i++ {
			digit := int(x[i] - '0')
			if limit < digit {
				count += int64(math.Pow(float64(limit+1), float64(preLen-i)))
				return count
			}
			count += int64(digit) * int64(math.Pow(float64(limit+1), float64(preLen-1-i)))
		}

		if suffix >= s {
			count += 1
		}

		return count
	}

	startStr := strconv.FormatInt(start-1, 10)
	finishStr := strconv.FormatInt(finish, 10)

	return calculate(finishStr) - calculate(startStr)
}
