package main

import "sort"

func frequencySort(nums []int) []int {
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}

	sort.SliceStable(nums, func(i, j int) bool {
		if countMap[nums[i]] == countMap[nums[j]] {
			return nums[i] > nums[j]
		}
		return countMap[nums[i]] < countMap[nums[j]]
	})

	return nums
}
