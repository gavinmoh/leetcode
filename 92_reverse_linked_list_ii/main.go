package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	currentNode := head
	index := 1
	subNodes := make(map[int]*ListNode)

	var nodeBefore *ListNode
	var nodeAfter *ListNode
	var reversedHead *ListNode
	var reversedTail *ListNode
	for currentNode != nil && index <= right {
		if index == left-1 {
			nodeBefore = currentNode
		}

		if index == left {
			reversedTail = currentNode

			// break the link between the node before and the node to be reversed
			if nodeBefore != nil {
				nodeBefore.Next = nil
			}
		}

		if index == right {
			nodeAfter = currentNode.Next
			reversedHead = currentNode

			// break the link between the node to be reversed and the node after
			if nodeAfter != nil {
				currentNode.Next = nil
			}
		}

		if index >= left && index <= right {
			subNodes[index] = currentNode
		}

		currentNode = currentNode.Next
		index++
	}

	// reverse the nodes
	for i := right; i >= left; i-- {
		subNodes[i].Next = subNodes[i-1]
	}

	// relink the head and tail of the reversed nodes
	if nodeAfter != nil {
		reversedTail.Next = nodeAfter
	}

	if nodeBefore != nil {
		nodeBefore.Next = reversedHead
	} else {
		return reversedHead
	}

	return head
}
