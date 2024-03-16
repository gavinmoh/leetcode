package main

func findMaxLength(nums []int) int {
	runningSum := 0
	maxLength := 0
	first := make(map[int]int)
	first[0] = -1

	for right, num := range nums {
		if num == 0 {
			runningSum--
		} else {
			runningSum++
		}
		if left, ok := first[runningSum]; !ok {
			first[runningSum] = right
		} else {
			length := right - left
			maxLength = max(length, maxLength)
		}
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
