package main

import "math"

func countTriples(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			sqrt := math.Sqrt(float64(i*i + j*j))
			intPart, fracPart := math.Modf(sqrt)
			if fracPart != 0 {
				continue
			}

			if int(intPart) <= n {
				count++
			}
		}
	}

	return count * 2
}
