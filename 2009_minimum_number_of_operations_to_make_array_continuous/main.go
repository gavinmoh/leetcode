package main

import "sort"

func minOperations(nums []int) int {
	length := len(nums)
	sort.Ints(nums)

	// remove duplicates
	index := 1
	for i := 1; i < length; i++ {
		if nums[i] != nums[i-1] {
			nums[index] = nums[i]
			index++
		}
	}

	count := length
	for i, j := 0, 0; i < index; i++ {
		for j < index && nums[j]-nums[i] <= length-1 {
			j++
		}
		count = min(count, length-(j-i))
	}

	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
