package main

func countPartitions(nums []int) int {
	n := len(nums)
	prefixSum := make([]int, n)
	prefixSum[0] = nums[0]

	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}

	count := 0
	for i := 0; i < n-1; i++ {
		left, right := prefixSum[i], prefixSum[n-1]-prefixSum[i]
		diff := left - right
		if diff%2 == 0 {
			count++
		}
	}

	return count
}
