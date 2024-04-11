package main

func removeKdigits(num string, k int) string {
	stack := []rune{}
	for _, char := range num {
		for len(stack) > 0 && k > 0 && stack[len(stack)-1] > char {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, char)
	}

	// remove remaing from back
	stack = stack[:len(stack)-k]

	// remove leading zeroes
	for len(stack) > 1 && stack[0] == '0' {
		stack = stack[1:]
	}

	if len(stack) == 0 {
		return "0"
	}

	return string(stack)
}
