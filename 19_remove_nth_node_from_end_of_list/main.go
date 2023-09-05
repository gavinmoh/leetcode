package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	currentNode := head
	length := 0

	nodes := make(map[int]*ListNode) // key: index, value: node

	for currentNode != nil {
		length++
		nodes[length] = currentNode
		currentNode = currentNode.Next
	}

	if length == 1 {
		return nil
	}

	position := (length + 1) - n
	target := nodes[position]
	if position == 1 { // remove the first node
		head = nodes[2]
	} else if position == length { // remove the last node
		previousNode := nodes[position-1]
		previousNode.Next = nil
	} else { // remove the middle node
		previousNode := nodes[position-1]
		previousNode.Next = target.Next
	}

	return head
}
