package main

func decrypt(code []int, k int) []int {
	n := len(code)
	circular := append(append(code, code...), code...)

	result := []int{}
	for i := 0; i < n; i++ {
		sum := 0

		if k > 0 {
			for j := 1; j <= k; j++ {
				sum += circular[n+i+j]
			}
		} else if k < 0 {
			for j := -1; j >= k; j-- {
				sum += circular[n+i+j]
			}
		}

		result = append(result, sum)
	}

	return result
}
