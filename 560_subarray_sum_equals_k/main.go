package main

func subarraySum(nums []int, k int) int {
	prefixSum := make(map[int]int) // prefixSum -> count
	prefixSum[0] = 1
	result := 0
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		diff := sum - k
		if count, ok := prefixSum[diff]; ok {
			result += count
		}
		prefixSum[sum]++
	}

	return result
}
