package main

func sumZero(n int) []int {
	result := make([]int, n)

	for i := 1; i < n; i += 2 {
		result[i-1] = -i
		result[i] = i
	}

	return result
}
