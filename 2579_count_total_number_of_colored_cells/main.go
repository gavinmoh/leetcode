package main

func coloredCells(n int) int64 {
	return IntPow(int64(n), int64(2)) + IntPow(int64(n-1), int64(2))
}

func IntPow(base, exp int64) int64 {
	result := int64(1)
	for {
		if exp&1 == 1 {
			result *= base
		}
		exp >>= 1
		if exp == 0 {
			break
		}
		base *= base
	}

	return result
}
