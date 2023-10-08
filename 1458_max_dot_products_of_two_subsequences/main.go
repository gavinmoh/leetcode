package main

func maxDotProduct(nums1 []int, nums2 []int) int {
	nums1Size := len(nums1)
	nums2Size := len(nums2)

	dp := make([][]int, nums1Size)
	for i := range dp {
		dp[i] = make([]int, nums2Size)
	}

	for i := 0; i < nums1Size; i++ {
		for j := 0; j < nums2Size; j++ {
			dp[i][j] = nums1[i] * nums2[j]
			if i > 0 && j > 0 {
				dp[i][j] += max(dp[i-1][j-1], 0)
			}

			if i > 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j])
			}

			if j > 0 {
				dp[i][j] = max(dp[i][j], dp[i][j-1])
			}
		}
	}

	return dp[nums1Size-1][nums2Size-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
