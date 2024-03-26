package main

func firstMissingPositive(nums []int) int {
	n := len(nums)
	seen := make(map[int]bool, n)
	for _, num := range nums {
		seen[num] = true
	}

	for i := 1; i <= n; i++ {
		if !seen[i] {
			return i
		}
	}

	return n + 1
}
