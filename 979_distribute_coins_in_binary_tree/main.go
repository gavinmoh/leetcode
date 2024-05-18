package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distributeCoins(root *TreeNode) int {
	count := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		count += abs(left) + abs(right)

		return (node.Val - 1) + left + right
	}
	dfs(root)

	return count
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
