package main

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make(map[int]map[int]int)

	var dfs func(int, int) int
	dfs = func(i int, j int) int {
		if i >= m || j >= n {
			return 0
		}

		if _, ok := dp[i]; !ok {
			dp[i] = make(map[int]int)
		}

		if cachedResult, ok := dp[i][j]; ok {
			return cachedResult
		}

		if text1[i] == text2[j] {
			dp[i][j] = 1 + dfs(i+1, j+1)
		} else {
			dp[i][j] = max(dfs(i+1, j), dfs(i, j+1))
		}
		return dp[i][j]
	}

	return dfs(0, 0)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
