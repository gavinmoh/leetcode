package main

func countSubarrays(nums []int) int {
	n := len(nums)
	result := 0

	for i := 0; i < n-2; i++ {
		sum := nums[i] + nums[i+2]
		if sum*2 == nums[i+1] {
			result++
		}
	}

	return result
}
