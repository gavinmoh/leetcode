package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {

	var dfs func(*TreeNode) []int
	dfs = func(node *TreeNode) []int {
		result := []int{}

		if node == nil {
			return result
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		if node.Left == nil && node.Right == nil {
			result = append(result, node.Val)
		}

		result = append(result, left...)
		result = append(result, right...)

		return result
	}

	root1Vals := dfs(root1)
	root2Vals := dfs(root2)

	if len(root1Vals) != len(root2Vals) {
		return false
	}

	for i, val := range root1Vals {
		if val != root2Vals[i] {
			return false
		}
	}

	return true
}
