package main

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		dp[i] = n
	}
	dp[0] = 0

	for target := 1; target < n+1; target++ {
		for s := 1; s < target+1; s++ {
			square := s * s
			if target-square < 0 {
				break
			}
			dp[target] = min(dp[target], 1+dp[target-square])
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
