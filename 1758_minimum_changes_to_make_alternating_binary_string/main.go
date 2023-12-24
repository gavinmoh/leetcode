package main

func minOperations(s string) int {
	startWithOne := 0
	startWithZero := 0

	for i, digit := range s {
		if (i+1)%2 == 0 {
			if digit == '0' {
				startWithOne++
			} else {
				startWithZero++
			}
		} else {
			if digit == '1' {
				startWithOne++
			} else {
				startWithZero++
			}
		}
	}

	if startWithOne == len(s) || startWithZero == len(s) {
		return 0
	}

	return min(startWithOne, startWithZero)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
