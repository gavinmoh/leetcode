package main

func maxAbsoluteSum(nums []int) int {
	// kadane's algorithm
	maxSum, minSum := nums[0], nums[0]
	currentMax, currentMin := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		currentMax = max(currentMax+nums[i], nums[i])
		maxSum = max(maxSum, currentMax)

		currentMin = min(currentMin+nums[i], nums[i])
		minSum = min(minSum, currentMin)
	}

	return max(abs(maxSum), abs(minSum))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
