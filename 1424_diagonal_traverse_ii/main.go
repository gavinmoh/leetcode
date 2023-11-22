package main

import (
	"sort"
)

func findDiagonalOrder(nums [][]int) []int {

	diagonal := [][]int{}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			// sum, row, val
			diagonal = append(diagonal, []int{i + j, i, nums[i][j]})
		}
	}

	// sort by sum asc, row desc
	sort.Slice(diagonal, func(i, j int) bool {
		if diagonal[i][0] == diagonal[j][0] {
			return diagonal[i][1] > diagonal[j][1]
		}
		return diagonal[i][0] < diagonal[j][0]
	})

	output := make([]int, len(diagonal))
	for i := 0; i < len(diagonal); i++ {
		output[i] = diagonal[i][2]
	}
	return output
}
