package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	length := 0
	for current := head; current != nil; {
		length++
		current = current.Next
	}

	mid := (length / 2) + 1
	for i := 1; i <= length; i++ {
		if i == mid {
			break
		}
		head = head.Next
	}

	return head
}
