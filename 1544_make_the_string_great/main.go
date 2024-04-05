package main

func makeGood(s string) string {
	stack := []rune{}
	for _, char := range s {
		stack = append(stack, char)

		n := len(stack)
		if n < 2 {
			continue
		}

		if abs(stack[n-1]-stack[n-2]) == 32 {
			stack = stack[:n-2]
		}
	}

	return string(stack)
}

func abs(x rune) rune {
	if x < 0 {
		return -x
	}
	return x
}
