package main

func maximumDifference(nums []int) int {
	maxDiff := -1
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				diff := abs(nums[i] - nums[j])
				if diff > maxDiff {
					maxDiff = diff
				}
			}
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
