package main

const MOD = 1e9 + 7

func numTilings(n int) int {
	dp := make([]int, max(3, n+1))
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = ((dp[i-1] * 2) + dp[i-3]) % MOD
	}

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
