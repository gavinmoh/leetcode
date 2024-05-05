package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	current := node
	var prev *ListNode
	for current != nil {
		if current.Next != nil {
			current.Val = current.Next.Val
		} else {
			prev.Next = nil
		}
		prev = current
		current = current.Next
	}
}
