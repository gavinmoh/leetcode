package main

func maxAdjacentDistance(nums []int) int {
	n := len(nums)
	maxDiff := abs(nums[0] - nums[n-1])
	for i := 0; i < n-1; i++ {
		diff := abs(nums[i] - nums[i+1])
		if diff > maxDiff {
			maxDiff = diff
		}
	}

	return maxDiff
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
