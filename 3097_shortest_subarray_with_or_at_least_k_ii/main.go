package main

import "math"

func minimumSubarrayLength(nums []int, k int) int {
	minLength := math.MaxInt
	bitCounts := make([]int, 32)

	updateBitCounts := func(num int, delta int) {
		for i := 0; i < 32; i++ {
			if num&(1<<i) != 0 {
				bitCounts[i] += delta
			}
		}
	}

	bitCountsToNum := func() int {
		result := 0
		for i := 0; i < 32; i++ {
			if bitCounts[i] != 0 {
				result |= (1 << i)
			}
		}
		return result
	}

	for left, right := 0, 0; right < len(nums); right++ {
		updateBitCounts(nums[right], 1)

		for left <= right && bitCountsToNum() >= k {
			minLength = min(minLength, right-left+1)
			updateBitCounts(nums[left], -1)
			left++
		}
	}

	if minLength == math.MaxInt {
		return -1
	}

	return minLength
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
