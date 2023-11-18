package main

import "sort"

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)

	left := 0
	currentSum := 0

	for right := 0; right < len(nums); right++ {
		target := nums[right]
		currentSum += target
		windowSize := right - left + 1

		// checking the number of operation needed
		// in order to increment all the element inside
		// the window to target, if yes, we slide the
		// window to the right (left + 1)
		if windowSize*target-currentSum > k {
			currentSum -= nums[left]
			left++
		}
	}

	return len(nums) - left
}
