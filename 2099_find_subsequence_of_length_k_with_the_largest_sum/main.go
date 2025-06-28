package main

import "sort"

func maxSubsequence(nums []int, k int) []int {
	sorted := make([][2]int, len(nums))
	for i, num := range nums {
		sorted[i] = [2]int{i, num}
	}

	sort.SliceStable(sorted, func(i, j int) bool {
		return sorted[i][1] < sorted[j][1]
	})

	subsequence := make([][2]int, k)
	for i := 0; i < k; i++ {
		subsequence[i] = sorted[len(nums)-1-i]
	}

	sort.SliceStable(subsequence, func(i, j int) bool {
		return subsequence[i][0] < subsequence[j][0]
	})

	result := []int{}
	for _, pair := range subsequence {
		result = append(result, pair[1])
	}

	return result
}
