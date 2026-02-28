package main

const MODULO = 1_000_000_007

func concatenatedBinary(n int) int {
	result := 0
	bits := 1
	x := 2

	for i := 1; i <= n; i++ {
		for x <= i {
			bits++
			x *= 2
		}

		result = (result << bits) + i
		result %= MODULO
	}

	return result
}
