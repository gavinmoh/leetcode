package main

func maxSubarrayLength(nums []int, k int) int {
	left := 0
	freq := make(map[int]int)
	maxLength := 0
	for right, num := range nums {
		freq[num]++
		for freq[num] > k {
			freq[nums[left]]--
			left++
		}
		length := right - left + 1
		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}
