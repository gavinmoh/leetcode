package main

// hint:
// 1. use dynamic programming
// 2. Let dp[i] be the solution for the prefix of the array that ends at index i, if the element at index i is in the subsequence.
// 3. dp[i] = nums[i] + max(0, dp[i-k], dp[i-k+1], ..., dp[i-1])
// 4. Use a heap with the sliding window technique to optimize the dp.
func constrainedSubsetSum(nums []int, k int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxSum := dp[0]
	heap := []int{dp[0]}
	for i := 1; i < len(nums); i++ {
		if i > k && heap[0] == dp[i-k-1] {
			heap = heap[1:]
		}
		dp[i] = nums[i] + max(0, heap[0])
		maxSum = max(maxSum, dp[i])
		for len(heap) > 0 && heap[len(heap)-1] < dp[i] {
			heap = heap[:len(heap)-1]
		}
		heap = append(heap, dp[i])
	}
	return maxSum

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
