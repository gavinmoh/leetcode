package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node.Left != nil {
			if node.Left.Left == nil && node.Left.Right == nil {
				sum += node.Left.Val
			}
			dfs(node.Left)
		}
		if node.Right != nil {
			dfs(node.Right)
		}
	}
	dfs(root)

	return sum
}
