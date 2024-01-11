package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	// node -> decendants
	nodes := make(map[int][]int)

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node.Left == nil && node.Right == nil {
			nodes[node.Val] = []int{node.Val}
			return
		}

		nodes[node.Val] = []int{node.Val}
		if node.Left != nil {
			dfs(node.Left)
			nodes[node.Val] = append(nodes[node.Val], nodes[node.Left.Val]...)
		}

		if node.Right != nil {
			dfs(node.Right)
			nodes[node.Val] = append(nodes[node.Val], nodes[node.Right.Val]...)
		}
	}
	dfs(root)

	maxDiff := math.MinInt
	for node, descendants := range nodes {
		for _, descendant := range descendants {
			maxDiff = max(maxDiff, abs(node-descendant))
		}
	}

	return maxDiff
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
