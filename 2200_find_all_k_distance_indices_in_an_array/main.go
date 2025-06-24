package main

import "sort"

func findKDistantIndices(nums []int, key int, k int) []int {
	result := []int{}
	keyIdx := []int{}
	for i, num := range nums {
		if num == key {
			keyIdx = append(keyIdx, i)
		}
	}

	for i := range nums {
		for _, j := range keyIdx {
			if abs(i-j) <= k {
				result = append(result, i)
				break
			}
		}
	}

	sort.Ints(result)

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
