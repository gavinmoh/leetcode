package main

var PRIMES = map[int]struct{}{
	2:  {},
	3:  {},
	5:  {},
	7:  {},
	11: {},
	13: {},
	17: {},
	19: {},
	23: {},
	29: {},
	31: {},
}

func countPrimeSetBits(left int, right int) int {
	count := 0

	for x := left; x <= right; x++ {
		setBits := countSetBits(x)
		if _, ok := PRIMES[setBits]; ok {
			count++
		}
	}

	return count
}

func countSetBits(n int) int {
	count := 0

	for n > 0 {
		if n&1 == 1 {
			count++
		}
		n >>= 1
	}

	return count
}
