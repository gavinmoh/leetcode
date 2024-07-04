package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	previous, current := head, head.Next
	for current.Next != nil {
		if current.Val == 0 {
			previous.Next = current
			previous = current
		}
		previous.Val += current.Val
		current = current.Next
	}
	previous.Next = nil

	return head
}
