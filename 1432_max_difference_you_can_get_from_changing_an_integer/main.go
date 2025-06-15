package main

func maxDiff(num int) int {
	digits := []int{}
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}

	n := len(digits)
	firstNonNineDigit := -1
	firstMoreThanOneDigit := -1

	for i := n - 1; i >= 0; i-- {
		if digits[i] != 9 {
			firstNonNineDigit = digits[i]
			break
		}
	}

	// find first digit more than 1
	if digits[n-1] == 1 {
		for i := n - 1; i >= 0; i-- {
			if digits[i] > 1 {
				firstMoreThanOneDigit = digits[i]
				break
			}
		}
	}

	maxNum, minNum := 0, 0
	for i := n - 1; i >= 0; i-- {
		digit := digits[i]
		maxNum *= 10
		minNum *= 10

		if digit == firstNonNineDigit {
			maxNum += 9
		} else {
			maxNum += digit
		}

		if firstMoreThanOneDigit == -1 {
			// replace the first one with 1
			if digit == digits[n-1] {
				minNum += 1
			} else {
				minNum += digit
			}
		} else {
			if firstMoreThanOneDigit == digit {
				minNum += 0
			} else {
				minNum += digit
			}
		}
	}

	return maxNum - minNum
}
