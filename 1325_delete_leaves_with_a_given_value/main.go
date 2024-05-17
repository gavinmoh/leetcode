package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		if node.Left != nil {
			dfs(node.Left)
			if isLeafNode(node.Left) && node.Left.Val == target {
				node.Left = nil
			}
		}

		if node.Right != nil {
			dfs(node.Right)
			if isLeafNode(node.Right) && node.Right.Val == target {
				node.Right = nil
			}
		}
	}

	root = &TreeNode{Val: -1, Left: root}
	dfs(root)

	return root.Left
}

func isLeafNode(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}
