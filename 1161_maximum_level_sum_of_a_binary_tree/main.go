package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	levelSum := []int{0}

	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if len(levelSum) < level+1 {
			levelSum = append(levelSum, 0)
		}
		levelSum[level] += node.Val

		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}

	dfs(root, 1)

	result := -1
	maxSum := math.MinInt
	for level, sum := range levelSum {
		if level == 0 {
			continue
		}

		if sum > maxSum {
			maxSum = sum
			result = level
		}
	}

	return result
}
