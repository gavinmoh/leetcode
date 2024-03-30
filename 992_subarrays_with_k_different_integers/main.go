package main

func subarraysWithKDistinct(nums []int, k int) int {
	return slidingWindowAtMost(nums, k) - slidingWindowAtMost(nums, k-1)
}

func slidingWindowAtMost(nums []int, distinctK int) int {
	left := 0
	freq := make(map[int]int)
	totalCount := 0

	for right, num := range nums {
		freq[num]++

		for len(freq) > distinctK {
			freq[nums[left]]--
			if freq[nums[left]] == 0 {
				delete(freq, nums[left])
			}
			left++
		}

		totalCount += right - left + 1
	}
	return totalCount
}
