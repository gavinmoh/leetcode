package main

func checkValidString(s string) bool {
	openBrackets, asterisks := Stack{}, Stack{}
	for i, char := range s {
		if char == '(' {
			openBrackets.Push(i)
		} else if char == '*' {
			asterisks.Push(i)
		} else {
			if openBrackets.Any() {
				openBrackets.Pop()
			} else if asterisks.Any() {
				asterisks.Pop()
			} else {
				return false
			}
		}
	}

	for openBrackets.Any() && asterisks.Any() {
		if openBrackets.Pop() > asterisks.Pop() {
			return false
		}
	}

	return openBrackets.IsEmpty()
}

type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Size() int {
	arr := *s

	return len(arr)
}

func (s *Stack) Pop() int {
	arr := *s
	n := len(arr)
	popped := arr[n-1]
	*s = arr[:n-1]

	return popped
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack) Any() bool {
	return s.Size() > 0
}
