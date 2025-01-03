package main

func waysToSplitArray(nums []int) int {
	sum := 0
	prefixSum := make([]int, len(nums)+1)
	for i, num := range nums {
		sum += num
		prefixSum[i+1] = sum
	}

	count := 0
	for i := 0; i < len(nums)-1; i++ {
		left := prefixSum[i+1]
		right := sum - prefixSum[i+1]

		if left >= right {
			count++
		}
	}

	return count
}
