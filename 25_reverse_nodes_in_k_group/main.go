package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	nodes := make(map[int]*ListNode)
	index := 0

	current := head
	for current != nil {
		nodes[index] = current
		current = current.Next
		index++
	}

	for i := k - 1; i < index; i += k {
		reverseNodes := []*ListNode{}
		for j := i; j > i-k; j-- {
			reverseNodes = append(reverseNodes, nodes[j])
		}
		for j, node := range reverseNodes {
			nodes[i-(k-1)+j] = node
		}
	}

	for i := 0; i < index-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	nodes[index-1].Next = nil

	return nodes[0]
}
