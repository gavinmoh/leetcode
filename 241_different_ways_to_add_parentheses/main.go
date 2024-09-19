package main

import "strconv"

func diffWaysToCompute(expression string) []int {
	var dfs func(int, int) []int
	dfs = func(left int, right int) []int {
		result := []int{}
		for i := left; i < right; i++ {
			if isOperator(expression[i]) {
				nums1 := dfs(left, i-1)
				nums2 := dfs(i+1, right)
				for _, num1 := range nums1 {
					for _, num2 := range nums2 {
						result = append(result, calculate(num1, num2, expression[i]))
					}
				}
			}
		}

		if len(result) == 0 {
			num, _ := strconv.Atoi(expression[left : right+1])
			result = append(result, num)
		}

		return result
	}

	return dfs(0, len(expression)-1)
}

func isOperator(char byte) bool {
	switch char {
	case '+', '-', '*':
		return true
	}

	return false
}

func calculate(num1 int, num2 int, operator byte) int {
	var result int
	switch operator {
	case '+':
		result = num1 + num2
	case '-':
		result = num1 - num2
	case '*':
		result = num1 * num2
	}

	return result
}
