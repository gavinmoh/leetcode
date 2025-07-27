package main

func countHillValley(nums []int) int {
	// remove duplicates from nums without creating new slice
	right := 1
	for left := 1; left < len(nums); left++ {
		if nums[left] == nums[left-1] {
			continue
		}

		nums[right] = nums[left]
		right++
	}

	count := 0
	for i := 1; i < right-1; i++ {
		if nums[i] == nums[i-1] || nums[i] == nums[i+1] {
			continue
		}

		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			count++
		} else if nums[i] < nums[i-1] && nums[i] < nums[i+1] {
			count++
		}
	}

	return count
}
