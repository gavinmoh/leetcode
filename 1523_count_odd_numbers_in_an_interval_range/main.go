package main

func countOdds(low int, high int) int {
	n := high - low + 1
	if n%2 == 0 {
		return n / 2
	}

	if high%2 != 0 || low%2 != 0 {
		return n/2 + 1
	}

	return n / 2
}
