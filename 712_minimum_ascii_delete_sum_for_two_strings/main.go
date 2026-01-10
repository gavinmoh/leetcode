package main

func minimumDeleteSum(s1 string, s2 string) int {
	dp := map[int]map[int]int{}

	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i >= len(s1) || j >= len(s2) {
			return 0
		}

		if _, ok := dp[i]; !ok {
			dp[i] = map[int]int{}
		}

		if _, ok := dp[i][j]; ok {
			return dp[i][j]
		}

		if s1[i] == s2[j] {
			dp[i][j] = int(s1[i]) + dfs(i+1, j+1)
		} else {
			dp[i][j] = max(dfs(i+1, j), dfs(i, j+1))
		}

		return dp[i][j]
	}

	totalSum := dfs(0, 0)
	s1Sum, s2Sum := 0, 0
	for _, c := range s1 {
		s1Sum += int(c)
	}
	for _, c := range s2 {
		s2Sum += int(c)
	}

	return s1Sum + s2Sum - (2 * totalSum)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
