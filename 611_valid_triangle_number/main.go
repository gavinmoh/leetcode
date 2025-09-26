package main

import "sort"

func triangleNumber(nums []int) int {
	n := len(nums)
	count := 0
	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			side := nums[i] + nums[j]
			subArray := nums[j+1:]
			// look for nums with value < side
			count += sort.Search(len(subArray), func(idx int) bool {
				return subArray[idx] >= side
			})
		}
	}

	return count
}
