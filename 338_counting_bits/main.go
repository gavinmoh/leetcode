package main

func CountBits(n int) []int {
	results := make([]int, n+1)

	for i := 0; i <= n; i++ {
		count := 0
		number := i
		for number != 0 {
			// use bitwise AND to check if the last bit is 1
			// e.g. 101 & 1 = 1; 100 & 1 = 0
			count += number & 1
			// use bitwise right shift to remove the last bit
			// e.g. 101 >> 1 = 10
			number >>= 1
		}
		results[i] = count
	}

	return results
}
