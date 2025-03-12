package main

func maximumCount(nums []int) int {
	negCount := 0
	posCount := len(nums)

	for _, num := range nums {
		if num < 0 {
			negCount++
			posCount--
		} else if num == 0 {
			posCount--
		} else {
			break
		}
	}

	return max(negCount, posCount)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
