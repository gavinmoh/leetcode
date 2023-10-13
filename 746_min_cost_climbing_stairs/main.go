package main

func minCostClimbingStairs(cost []int) int {
	length := len(cost)
	dp := make([]int, length+1)

	for i := 2; i <= length; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[length]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
