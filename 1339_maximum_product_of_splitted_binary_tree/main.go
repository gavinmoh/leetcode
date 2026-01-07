package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const modulo = 1_000_000_007

func maxProduct(root *TreeNode) int {
	totalSum := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		totalSum += node.Val
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)

	maxProduct := 0
	var dfs2 func(node *TreeNode) int
	dfs2 = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftSum := dfs2(node.Left)
		rightSum := dfs2(node.Right)

		if node.Left != nil {
			maxProduct = max(maxProduct, (totalSum-leftSum)*leftSum)
		}

		if node.Right != nil {
			maxProduct = max(maxProduct, (totalSum-rightSum)*rightSum)
		}

		return node.Val + leftSum + rightSum
	}

	dfs2(root)

	return maxProduct % modulo
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
