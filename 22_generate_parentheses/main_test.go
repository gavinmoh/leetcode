package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateParenthesis(t *testing.T) {
	case1 := 3
	case1Expected := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	case1Actual := GenerateParenthesis(case1)
	// fmt.Println(case1Actual)
	assert.Equal(t, 5, len(case1Actual))
	for _, expected := range case1Expected {
		assert.Contains(t, GenerateParenthesis(case1), expected)
	}

	case2 := 1
	case2Expected := []string{"()"}
	case2Actual := GenerateParenthesis(case2)
	assert.Equal(t, 1, len(case2Actual))
	for _, expected := range case2Expected {
		assert.Contains(t, case2Actual, expected)
	}

	case3 := 4
	case3Expected := []string{
		"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())",
		"(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()",
		"()()(())", "()()()()",
	}
	case3Actual := GenerateParenthesis(case3)
	assert.Equal(t, 14, len(case3Actual))
	for _, expected := range case3Expected {
		assert.Contains(t, case3Actual, expected)
	}
}
