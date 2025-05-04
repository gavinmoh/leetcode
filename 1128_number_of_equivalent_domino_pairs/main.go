package main

func numEquivDominoPairs(dominoes [][]int) int {
	freq := make([]int, 100)
	result := 0

	for _, pair := range dominoes {
		a, b := pair[0], pair[1]
		if a > b {
			a, b = b, a
		}

		x := a*10 + b
		result += freq[x]
		freq[x]++
	}

	return result
}
