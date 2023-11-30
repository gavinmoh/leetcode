package main

import "math"

func minimumOneBitOperations(n int) int {
	if n == 0 {
		return 0
	}

	// find the most significant bit
	k := 0
	for int(math.Pow(2, float64(k))) <= n {
		k++
	}
	k--

	operation := int(math.Pow(2, float64(k+1)) - 1)
	remainder := int(math.Pow(2, float64(k))) ^ n

	return operation - minimumOneBitOperations(remainder)

}
