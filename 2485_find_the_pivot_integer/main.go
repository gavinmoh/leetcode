package main

func pivotInteger(n int) int {
	arithmeticSum := n * (1 + n) / 2
	runningSum := 0
	for i := 1; i <= n; i++ {
		if (arithmeticSum - runningSum) == (runningSum + i) {
			return i
		}
		runningSum += i
	}

	return -1
}
