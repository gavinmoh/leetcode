package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	nodes := make(map[int]*ListNode)
	index := 0

	var dfs func(*ListNode)
	dfs = func(node *ListNode) {
		if node == nil {
			return
		}
		nodes[index] = node
		index++
		dfs(node.Next)
	}
	dfs(head)

	for i := 0; i < index; i++ {
		if i%2 != 0 {
			nodes[i-1], nodes[i] = nodes[i], nodes[i-1]
		}
	}

	for i := 0; i < index-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	nodes[index-1].Next = nil

	return nodes[0]
}
