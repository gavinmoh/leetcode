package main

func shortestPalindrome(s string) string {
	position := len(s)
	if position <= 1 {
		return s
	}

	for right := len(s) - 1; right > 0; right-- {
		if isPalindrome(s[0 : right+1]) {
			break
		} else {
			position = right
		}
	}

	for _, char := range s[position:] {
		s = string(char) + s
	}

	return s
}

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
