package main

func countLargestGroup(n int) int {
	groups := map[int]int{}
	maxCount := 0
	for num := 1; num <= n; num++ {
		sum := digitSum(num)
		groups[sum]++
		maxCount = max(maxCount, groups[sum])
	}

	result := 0
	for _, count := range groups {
		if count == maxCount {
			result++
		}
	}

	return result
}

func digitSum(x int) int {
	sum := 0
	for x > 0 {
		sum += x % 10
		x /= 10
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
