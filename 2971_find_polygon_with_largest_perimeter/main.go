package main

import "sort"

func largestPerimeter(nums []int) int64 {
	sort.Ints(nums)
	n := len(nums)
	prefixSum := make([]int, n)

	currentSum := 0
	for i, num := range nums {
		currentSum += num
		prefixSum[i] = currentSum
	}

	left, right := 0, n-1
	for right-left >= 2 {
		longest := nums[right]
		shorterSides := prefixSum[right-1]
		if left != 0 {
			shorterSides -= prefixSum[left-1]
		}

		if shorterSides > longest {
			return int64(shorterSides + longest)
		}
		right--
	}

	return -1
}
