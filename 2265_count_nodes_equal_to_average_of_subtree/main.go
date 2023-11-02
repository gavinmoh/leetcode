package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	count := 0

	var dfs func(node *TreeNode) []int
	dfs = func(node *TreeNode) []int {
		if node == nil {
			return []int{0, 0}
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		sum := left[0] + right[0] + node.Val
		nodesCount := left[1] + right[1] + 1

		if sum/nodesCount == node.Val {
			count++
		}

		return []int{sum, nodesCount}
	}

	dfs(root)

	return count
}
