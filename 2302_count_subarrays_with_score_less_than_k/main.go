package main

func countSubarrays(nums []int, k int64) int64 {
	n := len(nums)
	result := int64(0)
	runningSum := int64(0)

	for left, right := 0, 0; right < n; right++ {
		runningSum += int64(nums[right])

		// shrink the window
		for left <= right && runningSum*int64(right-left+1) >= k {
			runningSum -= int64(nums[left])
			left++
		}

		result += int64(right - left + 1)
	}

	return result
}
