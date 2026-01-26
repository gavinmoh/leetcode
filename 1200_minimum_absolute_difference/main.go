package main

import (
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	result := [][]int{}
	minDiff := math.MaxInt

	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		if diff == minDiff {
			result = append(result, []int{arr[i-1], arr[i]})
		} else if diff < minDiff {
			minDiff = diff
			result = [][]int{{arr[i-1], arr[i]}}
		}
	}

	return result
}
