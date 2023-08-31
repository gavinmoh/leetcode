package main

func IsValid(s string) bool {
	validPair := func(a, b rune) bool {
		if a == '(' && b == ')' {
			return true
		}
		if a == '[' && b == ']' {
			return true
		}
		if a == '{' && b == '}' {
			return true
		}
		return false
	}
	isOpen := func(char rune) bool {
		if char == '(' || char == '[' || char == '{' {
			return true
		}
		return false
	}
	isClose := func(char rune) bool {
		if char == ')' || char == ']' || char == '}' {
			return true
		}
		return false
	}

	stack := []rune{}
	for _, char := range s {
		if len(stack) == 0 && isClose(char) {
			return false
		}

		if len(stack) == 0 {
			stack = append(stack, char)
			continue
		}

		if isOpen(char) {
			stack = append(stack, char)
			continue
		}

		top := stack[len(stack)-1]
		if validPair(top, char) {
			stack = stack[:len(stack)-1]
			continue
		}
		return false
	}
	return len(stack) == 0
}
