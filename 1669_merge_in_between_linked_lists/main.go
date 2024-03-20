package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	var left *ListNode
	var right *ListNode
	current := list1
	index := 0
	for current != nil {
		if index == a-1 {
			left = current
		} else if index == b+1 {
			right = current
			break
		}
		current = current.Next
		index++
	}

	left.Next = list2
	current = list2
	for current != nil {
		if current.Next == nil {
			current.Next = right
			break
		}
		current = current.Next
	}

	return list1
}
