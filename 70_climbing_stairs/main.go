package main

func climbStairs(n int) int {
	// fibonacci
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return b
}
