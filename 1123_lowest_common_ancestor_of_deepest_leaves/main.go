package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	maxLevel := 0
	lca := root

	var dfs func(node *TreeNode, level int) int
	dfs = func(node *TreeNode, level int) int {
		if level > maxLevel {
			maxLevel = level
		}

		leftDeepest, rightDeepest := level, level
		if node.Left != nil {
			leftDeepest = dfs(node.Left, level+1)
		}

		if node.Right != nil {
			rightDeepest = dfs(node.Right, level+1)
		}

		if leftDeepest == rightDeepest && leftDeepest == maxLevel {
			lca = node
		}

		if leftDeepest > rightDeepest {
			return leftDeepest
		}

		return rightDeepest
	}

	dfs(root, 0)

	return lca
}
