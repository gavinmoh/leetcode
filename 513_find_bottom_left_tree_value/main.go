package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	maxLevel := -1
	result := -1

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		dfs(node.Left, level+1)
		dfs(node.Right, level+1)

		if level > maxLevel {
			result = node.Val
			maxLevel = level
		}
	}

	dfs(root, 0)
	return result
}
