package main

func getNoZeroIntegers(n int) []int {
	for a := 1; a < n; a++ {
		b := n - a
		if isNoZeroInteger(a) && isNoZeroInteger(b) {
			return []int{a, b}
		}
	}

	return []int{-1, -1}
}

func isNoZeroInteger(x int) bool {
	for x > 0 {
		if x%10 == 0 {
			return false
		}
		x /= 10
	}

	return true
}
