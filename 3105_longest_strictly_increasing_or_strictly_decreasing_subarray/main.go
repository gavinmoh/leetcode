package main

func longestMonotonicSubarray(nums []int) int {
	maxIncreasing, maxDecreasing := 1, 1
	currentIncreasing, currentDecreasing := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			currentIncreasing++
			currentDecreasing = 1
		} else if nums[i] < nums[i-1] {
			currentIncreasing = 1
			currentDecreasing++
		} else {
			currentIncreasing = 1
			currentDecreasing = 1
		}

		maxIncreasing = max(maxIncreasing, currentIncreasing)
		maxDecreasing = max(maxDecreasing, currentDecreasing)
	}

	return max(maxIncreasing, maxDecreasing)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
