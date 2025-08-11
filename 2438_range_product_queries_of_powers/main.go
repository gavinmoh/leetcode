package main

func productQueries(n int, queries [][]int) []int {
	powers := []int{}
	for i := 0; n > 0; i++ {
		num := (n & 1) << i
		n >>= 1

		if num == 0 {
			continue
		}

		powers = append(powers, num)
	}

	mod := 1_000_000_007
	results := make([]int, len(queries))
	for i, query := range queries {
		left, right := query[0], query[1]
		product := 1
		for j := left; j <= right; j++ {
			product = (product * powers[j]) % mod
		}
		results[i] = product
	}

	return results
}
