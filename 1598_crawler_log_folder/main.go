package main

func minOperations(logs []string) int {
	stack := []string{}

	for _, log := range logs {
		if log == "./" {
			continue
		}

		if log == "../" {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, log)
		}
	}

	return len(stack)
}
