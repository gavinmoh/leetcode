package main

func maxScore(s string) int {
	maximumScore := -1
	for i := 1; i < len(s); i++ {
		leftStr := s[:i]
		rightStr := s[i:]

		left := 0
		right := 0
		// count zeroes on the left
		for _, digit := range leftStr {
			if digit == '0' {
				left++
			}
		}

		// count ones on the right
		for _, digit := range rightStr {
			if digit == '1' {
				right++
			}
		}

		score := left + right
		maximumScore = max(score, maximumScore)
	}

	return maximumScore
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
