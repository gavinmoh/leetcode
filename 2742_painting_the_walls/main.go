package main

func paintWalls(cost []int, time []int) int {
	dp := make([][]int, len(cost)+1)
	for i := range dp {
		dp[i] = make([]int, len(cost)+1)
	}

	var dfs func(i, remain int) int
	dfs = func(i, remain int) int {
		if remain <= 0 {
			return 0
		}

		if i == len(cost) {
			return 1 << 31 // max int
		}

		if dp[i][remain] != 0 {
			return dp[i][remain]
		}

		paint := cost[i] + dfs(i+1, remain-1-time[i])
		skip := dfs(i+1, remain)
		dp[i][remain] = min(paint, skip)
		return dp[i][remain]
	}

	return dfs(0, len(cost))

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
