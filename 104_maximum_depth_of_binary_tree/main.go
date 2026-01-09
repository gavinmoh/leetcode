package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	maxDepth := 0
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		if depth > maxDepth {
			maxDepth = depth
		}

		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}

	dfs(root, 1)

	return maxDepth
}
