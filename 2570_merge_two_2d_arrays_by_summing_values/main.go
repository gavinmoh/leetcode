package main

import "sort"

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	nums := map[int]int{}
	for _, num := range nums1 {
		id, val := num[0], num[1]
		nums[id] += val
	}

	for _, num := range nums2 {
		id, val := num[0], num[1]
		nums[id] += val
	}

	result := [][]int{}
	for id, val := range nums {
		result = append(result, []int{id, val})
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})

	return result
}
