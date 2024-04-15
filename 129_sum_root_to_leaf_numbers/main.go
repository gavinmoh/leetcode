package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	sum := 0

	var dfs func(node *TreeNode, acc int)
	dfs = func(node *TreeNode, acc int) {
		if node.Left == nil && node.Right == nil {
			sum += (acc * 10) + node.Val
			return
		}

		if node.Left != nil {
			dfs(node.Left, (acc*10)+node.Val)
		}
		if node.Right != nil {
			dfs(node.Right, (acc*10)+node.Val)
		}
	}
	dfs(root, 0)

	return sum
}
