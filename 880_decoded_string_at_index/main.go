package main

func decodeAtIndex(s string, k int) string {
	size := 0

	for _, char := range s {
		if isDigit(char) {
			size *= toInteger(char)
		} else {
			size++
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		k %= size
		char := rune(s[i])
		if k == 0 && !isDigit(char) {
			return string(char)
		}

		if isDigit(char) {
			size /= toInteger(char)
		} else {
			size--
		}
	}

	return ""
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func toInteger(char rune) int {
	return int(char - '0')
}
