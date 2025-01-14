package main

func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	freq := make([]int, n+1)
	result := make([]int, n)

	for i := 0; i < n; i++ {
		a, b := A[i], B[i]

		freq[a]++
		if freq[a] > 1 {
			result[i]++
		}

		freq[b]++
		if freq[b] > 1 {
			result[i]++
		}

		if i > 0 {
			result[i] += result[i-1]
		}
	}

	return result
}
