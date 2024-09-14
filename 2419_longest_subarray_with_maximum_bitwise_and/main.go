package main

func longestSubarray(nums []int) int {
	max, maxCount, currentCount := nums[0], 1, 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			if nums[i] == max {
				currentCount++
				if currentCount > maxCount {
					maxCount = currentCount
				}
			}
		} else if nums[i] > max {
			max, maxCount, currentCount = nums[i], 1, 1
		} else {
			currentCount = 1
		}
	}

	return maxCount
}
