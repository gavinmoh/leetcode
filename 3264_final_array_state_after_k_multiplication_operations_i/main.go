package main

import "math"

func getFinalState(nums []int, k int, multiplier int) []int {
	for i := 0; i < k; i++ {
		minValue := math.MaxInt
		minIndex := -1
		for idx, num := range nums {
			if num < minValue {
				minValue = num
				minIndex = idx
			}
		}
		nums[minIndex] *= multiplier
	}

	return nums
}
