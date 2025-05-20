package main

func isZeroArray(nums []int, queries [][]int) bool {
	n := len(nums)
	diffArray := make([]int, n+1)
	for _, query := range queries {
		left, right := query[0], query[1]
		diffArray[left] += 1
		diffArray[right+1] -= 1
	}

	prefixSum := make([]int, n)
	prefixSum[0] = diffArray[0]
	for i := 1; i < n; i++ {
		prefixSum[i] = diffArray[i] + prefixSum[i-1]
	}

	for i, num := range nums {
		if num-prefixSum[i] > 0 {
			return false
		}
	}

	return true
}
