package main

func maximum69Number(num int) int {
	digits := []int{}
	for num > 0 {
		digits = append(digits, num%10)
		num /= 10
	}

	// find the first 6 and change it to 9
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 6 {
			digits[i] = 9
			break
		}
	}

	result := 0
	for i := len(digits) - 1; i >= 0; i-- {
		result *= 10
		result += digits[i]
	}

	return result
}
