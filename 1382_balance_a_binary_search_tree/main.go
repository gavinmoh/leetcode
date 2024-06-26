package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	nums := []int{}

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		// inorder traversal
		dfs(node.Left)
		nums = append(nums, node.Val)
		dfs(node.Right)
	}
	dfs(root)

	// sort the slice
	sort.Ints(nums)

	var helper func(arr []int) *TreeNode
	helper = func(arr []int) *TreeNode {
		if len(arr) == 0 {
			return nil
		}

		mid := len(arr) / 2
		newRoot := &TreeNode{Val: arr[mid]}
		if len(arr) > 1 {
			newRoot.Left = helper(arr[:mid])
			newRoot.Right = helper(arr[mid+1:])
		}
		return newRoot
	}

	return helper(nums)
}
