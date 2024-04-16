package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil || level >= depth {
			return
		}

		if depth == level+1 {
			newLeft := &TreeNode{Val: val, Left: node.Left}
			node.Left = newLeft
			newRight := &TreeNode{Val: val, Right: node.Right}
			node.Right = newRight
			return
		}

		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}

	if depth == 1 {
		newRoot := &TreeNode{Val: val, Left: root}
		return newRoot
	}

	dfs(root, 1)
	return root
}
