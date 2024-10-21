package main

func maxUniqueSplit(s string) int {
	var dfs func(i int, currentHash map[string]bool) int
	dfs = func(i int, currentHash map[string]bool) int {
		if i == len(s) {
			return 0
		}

		result := 0
		for j := i; j < len(s); j++ {
			substr := s[i : j+1]

			if currentHash[substr] {
				continue
			}

			currentHash[substr] = true
			result = max(result, 1+dfs(j+1, currentHash))
			currentHash[substr] = false
		}

		return result
	}

	return dfs(0, map[string]bool{})
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
