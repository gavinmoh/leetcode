package main

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i, _ := range dp {
		dp[i] = 1
	}

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if nums[i] < nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	dpMax := 1
	for _, num := range dp {
		dpMax = max(num, dpMax)
	}
	return dpMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
