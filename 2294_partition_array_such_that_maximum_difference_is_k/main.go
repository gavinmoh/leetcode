package main

import "sort"

func partitionArray(nums []int, k int) int {
	count := 1
	sort.Ints(nums)

	currentMin := -1
	for _, num := range nums {
		if currentMin == -1 {
			currentMin = num
			continue
		}

		diff := num - currentMin
		if diff > k {
			currentMin = num
			count++
		}
	}

	return count
}
