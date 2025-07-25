package main

import "math"

func maxSum(nums []int) int {
	maxNum := math.MinInt
	positiveNums := map[int]bool{}

	for _, num := range nums {
		maxNum = max(maxNum, num)
		if num > 0 {
			positiveNums[num] = true
		}
	}

	if len(positiveNums) == 0 {
		return maxNum
	}

	sum := 0
	for num := range positiveNums {
		sum += num
	}

	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
