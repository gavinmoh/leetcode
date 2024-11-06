package main

import "math"

func getMaximumXor(nums []int, maximumBit int) []int {
	prefix := make([]int, len(nums))
	prefix[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] ^ nums[i]
	}

	result := []int{}
	e := int(math.Pow(2, float64(maximumBit)))
	for i := len(nums) - 1; i >= 0; i-- {
		result = append(result, prefix[i]^(e-1))
	}

	return result
}
