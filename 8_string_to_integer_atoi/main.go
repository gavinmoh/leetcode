package main

func myAtoi(s string) int {
	digits := ""
	substring := ""
	var sign int

	for _, char := range s {
		if char == ' ' && (len(substring) > 0 || sign != 0) {
			break
		}

		if char == ' ' && digits == "" {
			continue
		}

		if char == '-' && digits == "" {
			if substring != "" || sign != 0 {
				break
			}
			sign = -1
			continue
		}

		if char == '+' && digits == "" {
			if substring != "" || sign != 0 {
				break
			}
			sign = 1
			continue
		}

		if char >= '0' && char <= '9' {
			if substring != "" {
				break
			}
			digits += string(char)

			continue
		} else {
			if digits != "" {
				break
			}
		}

		substring += string(char)
	}

	if digits == "" {
		return 0
	}

	var result int
	for _, char := range digits {
		result = (result * 10) + int(char-'0')
		if result > 2147483648 || result < -2147483648 {
			break
		}
	}
	if sign != 0 {
		result = result * sign
	}

	// clamp result to 32-bit signed integer
	if result > 2147483647 {
		return 2147483647
	}

	if result < -2147483648 {
		return -2147483648
	}

	return result
}
