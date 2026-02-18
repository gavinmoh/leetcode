package main

func hasAlternatingBits(n int) bool {
	for n > 0 {
		if n>>1&1 == n&1 {
			return false
		}
		n >>= 1
	}

	return true
}
