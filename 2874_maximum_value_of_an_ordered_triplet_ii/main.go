package main

func maximumTripletValue(nums []int) int64 {
	n := len(nums)
	prefixMax, suffixMax := make([]int, n), make([]int, n)

	prefixMax[0] = nums[0]
	for i := 1; i < n; i++ {
		prefixMax[i] = max(nums[i], prefixMax[i-1])
	}

	suffixMax[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suffixMax[i] = max(nums[i], suffixMax[i+1])
	}

	maxValue := int64(0)
	for i := 1; i < n-1; i++ {
		value := int64(prefixMax[i-1]-nums[i]) * int64(suffixMax[i+1])
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
