package main

import "sort"

func minOperations(grid [][]int, x int) int {
	nums := []int{}
	for _, row := range grid {
		for _, num := range row {
			nums = append(nums, num)
		}
	}
	sort.Ints(nums)

	median := nums[len(nums)/2]
	count := 0

	for _, num := range nums {
		if abs(median-num)%x != 0 {
			return -1
		}
		count += abs(median-num) / x
	}

	return count
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
