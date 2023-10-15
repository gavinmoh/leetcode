package main

func numWays(steps int, arrLen int) int {
	mod := 1000000007

	dp := make([][]int, steps+1)
	for i := range dp {
		dp[i] = make([]int, steps)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(index, steps int) int {
		if steps == 0 {
			if index == 0 {
				return 1
			}
			return 0
		}
		if dp[steps][index] != -1 {
			return dp[steps][index]
		}

		// staying in the same place
		ans := dfs(index, steps-1)

		// moving left
		if index > 0 {
			ans += dfs(index-1, steps-1)
		}

		// moving right
		if index < arrLen-1 {
			ans += dfs(index+1, steps-1)
		}

		dp[steps][index] = ans % mod
		return dp[steps][index]
	}

	return dfs(0, steps)
}
