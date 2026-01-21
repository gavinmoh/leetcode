package main

import (
	"math"
)

func minBitwiseArray(nums []int) []int {
	result := make([]int, len(nums))
	for i, num := range nums {
		ans := -1
		for k := 0; k < num; k++ {
			x := num - int(math.Pow(2.0, float64(k)))
			if x < 0 {
				break
			}

			if x|(x+1) == num {
				ans = x
			}
		}

		result[i] = ans
	}

	return result
}
