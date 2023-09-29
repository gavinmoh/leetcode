package main

func isMonotonic(nums []int) bool {
	// edge cases
	if len(nums) <= 2 {
		return true
	}

	// find first non-equal pair in order to determine if increasing or decreasing
	index := 1
	for index < len(nums) && nums[index] == nums[index-1] {
		index++
	}

	increasing := true
	if index < len(nums) {
		increasing = nums[index] > nums[index-1]
	}

	// check if monotonic
	for index < len(nums) {
		if increasing && nums[index] < nums[index-1] {
			return false
		} else if !increasing && nums[index] > nums[index-1] {
			return false
		}
		index++
	}

	return true

}
