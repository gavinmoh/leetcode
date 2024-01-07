package main

func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	result := 0

	// dp[i][diff] is number of subsequence ending at i with diff
	dp := make(map[int]map[int]int)

	for i := 0; i < n; i++ {
		dp[i] = make(map[int]int)
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]
			dp[i][diff] += 1 + dp[j][diff]
			result += 1 + dp[j][diff]
		}
	}

	return result - (n * (n - 1) / 2)
}
