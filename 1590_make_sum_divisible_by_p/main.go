package main

func minSubarray(nums []int, p int) int {
	n := len(nums)
	totalSum := 0
	for _, num := range nums {
		totalSum = (totalSum + num) % p
	}

	target := totalSum % p
	if target == 0 {
		return 0
	}

	modMap := map[int]int{0: -1}
	currentSum := 0
	minLength := n

	for i, num := range nums {
		currentSum = (currentSum + num) % p
		needed := (currentSum - target + p) % p
		if index, ok := modMap[needed]; ok {
			minLength = min(minLength, i-index)
		}
		modMap[currentSum] = i
	}

	if minLength == n {
		return -1
	}

	return minLength
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
