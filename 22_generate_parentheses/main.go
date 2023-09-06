package main

func GenerateParenthesis(n int) []string {
	result := []string{}
	dfsParenthesis(n, n, &result, "")
	return result
}

func dfsParenthesis(left int, right int, result *[]string, s string) {
	// if left and right are both 0
	// that means we have exhausted all the parenthesis
	// therefore we can return the string
	if left == 0 && right == 0 {
		*result = append(*result, s)
		return
	}
	// if left is greater than 0, then we add a open parenthesis
	// and recursively call the function with left-1
	if left > 0 {
		dfsParenthesis(left-1, right, result, s+"(")
	}
	// if right is greater than 0, then we add a close parenthesis
	// and recursively call the function with right-1
	if right > 0 && left < right {
		dfsParenthesis(left, right-1, result, s+")")
	}
}
