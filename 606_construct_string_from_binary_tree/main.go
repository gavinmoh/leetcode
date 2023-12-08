package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	var dfs func(root *TreeNode) string
	dfs = func(root *TreeNode) string {
		if root.Left != nil {
			if root.Right != nil {
				return fmt.Sprintf("%d(%s)(%s)", root.Val, dfs(root.Left), dfs(root.Right))
			} else {
				return fmt.Sprintf("%d(%s)", root.Val, dfs(root.Left))
			}
		}
		if root.Right != nil {
			return fmt.Sprintf("%d()(%s)", root.Val, dfs(root.Right))
		}
		return fmt.Sprintf("%d", root.Val)
	}

	if root == nil {
		return ""
	}

	return dfs(root)
}
