package main

func rob(nums []int) int {
	// fibonacci
	// [rob1, rob2, n, n+1, ...]
	rob1, rob2 := 0, 0
	for _, money := range nums {
		rob1, rob2 = rob2, max(rob1+money, rob2)
	}
	return rob2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
