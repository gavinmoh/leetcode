package main

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	result := make([]int, n)
	stack := []int{}
	for i, temp := range temperatures {
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			// popping from stack
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[j] = abs(i - j)
		}
		stack = append(stack, i)
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
