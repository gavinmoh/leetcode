package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}

	var dfs func(i int, total int, current []int)
	dfs = func(i int, total int, current []int) {
		if total == target {
			newCurrent := make([]int, len(current))
			copy(newCurrent, current)
			result = append(result, newCurrent)
			return
		}

		if i >= len(candidates) || total > target {
			return
		}

		dfs(i+1, total+candidates[i], append(current, candidates[i]))
		for i < len(candidates)-1 && candidates[i] == candidates[i+1] {
			i++
		}
		dfs(i+1, total, current)
	}

	dfs(0, 0, []int{})

	return result
}
