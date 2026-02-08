package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	heights := map[*TreeNode]int{}

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left, right := dfs(node.Left), dfs(node.Right)
		heights[node] = max(left, right) + 1
		return heights[node]
	}

	dfs(root)

	for node := range heights {
		left, right := heights[node.Left], heights[node.Right]
		if abs(left-right) > 1 {
			return false
		}
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
