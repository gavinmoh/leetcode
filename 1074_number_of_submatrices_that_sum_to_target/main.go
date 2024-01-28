package main

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	m := len(matrix)
	n := len(matrix[0])
	count := 0

	for i := 0; i < m; i++ {
		nums := make([]int, n)
		for j := i; j < m; j++ {
			for k := 0; k < n; k++ {
				nums[k] += matrix[j][k]
			}
			count += subarraySum(nums, target)
		}
	}

	return count
}

func subarraySum(nums []int, k int) int {
	prefixSum := map[int]int{0: 1} // prefixSum -> count
	result := 0
	sum := 0

	for _, num := range nums {
		sum += num
		diff := sum - k
		if count, ok := prefixSum[diff]; ok {
			result += count
		}
		prefixSum[sum]++
	}

	return result
}
