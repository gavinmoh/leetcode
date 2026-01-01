package main

func plusOne(digits []int) []int {
	// increment by 1
	digits[len(digits)-1] += 1

	for i := len(digits) - 1; i > 0; i-- {
		if digits[i] > 9 {
			digits[i] -= 10
			digits[i-1] += 1
		} else {
			break
		}
	}

	if digits[0] > 9 {
		digits[0] -= 10
		digits = append([]int{1}, digits...)
	}

	return digits
}
