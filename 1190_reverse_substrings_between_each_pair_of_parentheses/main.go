package main

func reverseParentheses(s string) string {
	stack := []string{""}
	last := 0

	for _, letter := range s {
		switch letter {
		case '(':
			stack = append(stack, "")
			last++
		case ')':
			// reverse string of last string
			stack[last] = reverseString(stack[last])

			// merge with previous string
			if last > 0 {
				stack[last-1] += stack[last]
				stack = stack[:last]
				last--
			}
		default:
			stack[last] += string(letter)
		}
	}

	return stack[0]
}

func reverseString(s string) string {
	result := ""
	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}
