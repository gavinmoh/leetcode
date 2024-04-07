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

	filtered := make([]rune, len(stack)-count)
	j := len(stack) - count - 1
	for i := len(stack) - 1; i >= 0; i-- {
		if stack[i] == '(' && count > 0 {
			count--
			continue
		}

		filtered[j] = stack[i]
		j--
	}

	return string(filtered)
}

// func reverse(arr []rune) []rune {
// 	reversed := []rune{}
// 	for i := len(arr) - 1; i >= 0; i-- {
// 		reversed = append(reversed, arr[i])
// 	}
// 	return reversed
// }
