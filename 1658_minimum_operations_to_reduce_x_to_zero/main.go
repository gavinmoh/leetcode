package main

// this is a sliding window problem
// we want to find the longest subarray that sums to totalSum - x
// if we find it, then the answer is len(nums) - len(subarray)
func minOperations(nums []int, x int) int {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	target := totalSum - x
	currentSum := 0
	maxWindowSize := -1
	left := 0

	for right, num := range nums {
		currentSum += num
		for currentSum > target && left <= right {
			currentSum -= nums[left]
			left++
		}

		if currentSum == target {
			maxWindowSize = max(maxWindowSize, right-left+1)
		}
	}

	if maxWindowSize == -1 {
		return -1
	}

	return len(nums) - maxWindowSize
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
