package main

func countGoodNumbers(n int64) int {
	mod := int64(1_000_000_007)

	// fast exponentiation
	pow := func(x, y int64) int64 {
		result := int64(1)
		for y > 0 {
			if y%2 == 1 {
				result = result * x % mod
			}
			x = x * x % mod
			y /= 2
		}

		return result
	}

	return int(pow(5, (n+1)/2) * pow(4, n/2) % mod)
}
