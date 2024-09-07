package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	var dfs func(*ListNode, *TreeNode) bool
	dfs = func(node1 *ListNode, node2 *TreeNode) bool {
		if node1 == nil {
			return true
		}

		if node2 == nil {
			return false
		}

		if node1.Val != node2.Val {
			return false
		}

		if dfs(node1.Next, node2.Left) || dfs(node1.Next, node2.Right) {
			return true
		}

		return false
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		// pop
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			continue
		}

		if dfs(head, node) {
			return true
		}

		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}

	return false
}
