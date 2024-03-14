package main

func numSubarraysWithSum(nums []int, goal int) int {
	currentSum := 0
	prefixSum := make(map[int]int) // sum -> frequency
	count := 0

	for _, num := range nums {
		currentSum += num
		if currentSum == goal {
			count++
		}
		if freq, ok := prefixSum[currentSum-goal]; ok {
			count += freq
		}
		prefixSum[currentSum]++
	}

	return count
}
