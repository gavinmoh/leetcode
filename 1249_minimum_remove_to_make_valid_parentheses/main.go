package main

func minRemoveToMakeValid(s string) string {
	stack, count := []rune{}, 0
	for _, char := range s {
		if char == '(' {
			count++
			stack = append(stack, char)
		} else if char == ')' && count > 0 {
			count--
			stack = append(stack, char)
		} else if char != ')' {
			stack = append(stack, char)
		}
	}

	filtered := []rune{}
	for _, char := range reverse(stack) {
		if char == '(' && count > 0 {
			count--
			continue
		}

		filtered = append(filtered, char)
	}

	result := ""
	for _, char := range reverse(filtered) {
		result += string(char)
	}

	return result
}

func reverse(arr []rune) []rune {
	reversed := []rune{}
	for i := len(arr) - 1; i >= 0; i-- {
		reversed = append(reversed, arr[i])
	}
	return reversed
}
