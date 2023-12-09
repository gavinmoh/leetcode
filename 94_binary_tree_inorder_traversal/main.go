package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	output := []int{}

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		// first traverse the left
		if node.Left != nil {
			dfs(node.Left)
		}

		// operate on node itself
		output = append(output, node.Val)

		// then traverse the right
		if node.Right != nil {
			dfs(node.Right)
		}
	}

	dfs(root)
	return output
}
