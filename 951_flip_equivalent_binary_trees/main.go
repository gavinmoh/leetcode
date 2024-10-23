package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	var dfs func(node1 *TreeNode, node2 *TreeNode) bool
	dfs = func(node1 *TreeNode, node2 *TreeNode) bool {
		if node1 == nil && node2 == nil {
			return true
		}

		if node1 == nil || node2 == nil {
			return false
		}

		if node1.Val != node2.Val {
			return false
		}

		if isChildrenIdentical(node1, node2) || isChildrenIdentical(flipChildren(node1), node2) {
			return dfs(node1.Left, node2.Left) && dfs(node1.Right, node2.Right)
		}

		return false
	}

	return dfs(root1, root2)
}

func flipChildren(node *TreeNode) *TreeNode {
	node.Left, node.Right = node.Right, node.Left
	return node
}

func isChildrenIdentical(node1 *TreeNode, node2 *TreeNode) bool {
	node1Left, node1Right, node2Left, node2Right := -1, -1, -1, -1
	if node1.Left != nil {
		node1Left = node1.Left.Val
	}
	if node1.Right != nil {
		node1Right = node1.Right.Val
	}
	if node2.Left != nil {
		node2Left = node2.Left.Val
	}
	if node2.Right != nil {
		node2Right = node2.Right.Val
	}

	return node1Left == node2Left && node1Right == node2Right
}
