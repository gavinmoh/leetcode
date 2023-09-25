package main

func removeDuplicateLetters(s string) string {
	strMap := make(map[rune]int)
	stack := []rune{}
	stackMap := make(map[rune]bool)

	for _, char := range s {
		strMap[char]++
	}

	for i, char := range s {
		// first char
		if i == 0 {
			stack = append(stack, char)
			stackMap[char] = true
			strMap[char]--
			continue
		}

		// check if char is in stack
		// if yes, skip
		if presence, ok := stackMap[char]; ok && presence {
			strMap[char]--
			continue
		}

		for len(stack) > 0 {
			// compare char with last char in stack
			//  and check if char will occur again
			// if yes, pop last char in stack
			// if no, break and append char to stack
			stackTop := stack[len(stack)-1]
			if char < stackTop && strMap[stackTop] > 0 {
				stack = stack[:len(stack)-1]
				stackMap[stackTop] = false
				continue
			} else {
				break
			}
		}

		stack = append(stack, char)
		stackMap[char] = true
		strMap[char]--
	}

	return string(stack)
}
