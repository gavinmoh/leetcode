package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	sum := 0

	var dfs func(node *TreeNode, acc int)
	dfs = func(node *TreeNode, acc int) {
		if node == nil {
			return
		}

		acc = acc<<1 | node.Val
		if node.Left == nil && node.Right == nil {
			sum += acc
			return
		}

		dfs(node.Left, acc)
		dfs(node.Right, acc)
	}

	dfs(root, 0)
	return sum
}
