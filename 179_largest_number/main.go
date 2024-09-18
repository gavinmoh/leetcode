package main

import (
	"fmt"
	"sort"
)

func largestNumber(nums []int) string {
	sort.SliceStable(nums, func(i, j int) bool {
		left := fmt.Sprintf("%d%d", nums[i], nums[j])
		right := fmt.Sprintf("%d%d", nums[j], nums[i])
		return left > right
	})

	if nums[0] == 0 {
		return "0"
	}

	result := ""
	for _, num := range nums {
		result += fmt.Sprintf("%d", num)
	}

	return result
}
