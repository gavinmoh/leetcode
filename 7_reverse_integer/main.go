package main

func reverse(x int) int {
	var result int
	for x != 0 {
		remainder := x % 10
		result = (result * 10) + remainder
		x /= 10
	}

	// Check for overflow
	if result > 2147483647 || result < -2147483648 {
		return 0
	}

	return result
}
