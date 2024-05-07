package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func doubleIt(head *ListNode) *ListNode {
	head = &ListNode{Next: head}
	current, next := head, head.Next
	for current != nil && next != nil {
		next.Val *= 2
		current.Val += next.Val / 10
		next.Val %= 10

		current, next = current.Next, next.Next
	}

	if head.Val == 0 {
		head = head.Next
	}

	return head
}
