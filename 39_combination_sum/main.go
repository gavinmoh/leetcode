package main

func combinationSum(candidates []int, target int) [][]int {
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

		dfs(i, total+candidates[i], append(current, candidates[i]))
		dfs(i+1, total, current)
	}

	dfs(0, 0, []int{})

	return result
}
