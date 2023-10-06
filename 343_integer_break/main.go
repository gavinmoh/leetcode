package main

func integerBreak(n int) int {
	if n < 4 {
		return n - 1
	}

	product := 1
	for n > 4 {
		product *= 3
		n -= 3
	}

	return product * n
}
