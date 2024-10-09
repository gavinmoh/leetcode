package main

func minAddToMakeValid(s string) int {
	stack := []rune{}
	for _, char := range s {
		if len(stack) == 0 || char == '(' || stack[len(stack)-1] != '(' {
			stack = append(stack, char)
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack)
}
