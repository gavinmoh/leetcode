package main

func punishmentNumber(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		square := i * i
		if canPartition(square, i) {
			result += square
		}
	}

	return result
}

func canPartition(num int, target int) bool {
	if num == target {
		return true
	}

	for _, i := range []int{1_000_000, 100_000, 10_000, 1000, 100, 10} {
		x := num / i
		if x > 0 && canPartition(num%i, target-x) {
			return true
		}
	}

	return false
}
