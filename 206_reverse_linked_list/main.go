package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	nodes := []*ListNode{nil}
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}
	n := len(nodes)
	for i := n - 1; i > 0; i-- {
		nodes[i].Next = nodes[i-1]
	}

	return nodes[n-1]
}
