package main

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverFromPreorder(traversal string) *TreeNode {
	stack := [][]int{} // array of [level, val]
	current := []rune{}
	for _, char := range traversal {
		if char == '-' {
			if len(current) > 0 && current[len(current)-1] != '-' {
				stack = append(stack, strToStack(current))
				current = []rune{}
			}
		}
		current = append(current, char)
	}
	stack = append(stack, strToStack(current))

	var dfs func() *TreeNode
	dfs = func() *TreeNode {
		if len(stack) == 0 {
			return nil
		}

		level, node := stack[0][0], &TreeNode{Val: stack[0][1]}
		stack = stack[1:] // pop

		for len(stack) > 0 {
			nextLevel := stack[0][0]
			if level >= nextLevel {
				return node
			}

			if node.Left == nil {
				node.Left = dfs()
			} else {
				node.Right = dfs()
			}
		}

		return node
	}

	return dfs()
}

func strToStack(current []rune) []int {
	numStart := -1
	level := 0
	for i, char := range current {
		if char == '-' {
			level++
		} else if numStart == -1 {
			numStart = i
			break
		}
	}
	num, _ := strconv.Atoi(string(current[numStart:]))

	return []int{level, num}
}
