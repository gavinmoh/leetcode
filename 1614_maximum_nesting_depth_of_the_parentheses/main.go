package main

func maxDepth(s string) int {
	count, maxCount := 0, 0

	for _, char := range s {
		if char == '(' {
			count++
		}

		if char == ')' {
			count--
		}

		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}
