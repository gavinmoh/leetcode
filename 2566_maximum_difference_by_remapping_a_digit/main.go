package main

func minMaxDifference(num int) int {
	digits := []int{}
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}

	n := len(digits)
	firstNonNineDigit := -1
	firstNonZeroDigit := -1

	for i := n - 1; i >= 0; i-- {
		if digits[i] != 9 {
			firstNonNineDigit = digits[i]
			break
		}
	}

	for i := n - 1; i >= 0; i-- {
		if digits[i] != 0 {
			firstNonZeroDigit = digits[i]
			break
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

		if digit != firstNonZeroDigit {
			minNum += digit
		}
	}

	return maxNum - minNum
}
