package main

import "math"

func closestPrimes(left int, right int) []int {
	primes := make([]bool, right+1)
	for i := 2; i < len(primes); i++ {
		primes[i] = true
	}

	// Sieve of Eratosthenes
	for p := 2; p*p <= right; p++ {
		if primes[p] {
			for i := p * p; i < right+1; i += p {
				primes[i] = false
			}
		}
	}

	result := []int{-1, -1}
	prev, curr, dist := -1, -1, math.MaxInt
	for i := left; i <= right; i++ {
		if primes[i] {
			if curr == -1 {
				curr = i
				continue
			}

			prev = curr
			curr = i

			if curr-prev < dist {
				result[0] = prev
				result[1] = curr
				dist = curr - prev
			}
		}
	}

	return result
}
