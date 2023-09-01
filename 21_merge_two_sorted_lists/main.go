package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	var previousNode *ListNode
	currentNode := &ListNode{}
	head := currentNode

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			currentNode.Val = list1.Val
			list1 = list1.Next
		} else {
			currentNode.Val = list2.Val
			list2 = list2.Next
		}
		previousNode = currentNode
		currentNode.Next = &ListNode{}
		currentNode = currentNode.Next
	}

	for list1 != nil {
		currentNode.Val = list1.Val
		list1 = list1.Next
		previousNode = currentNode
		currentNode.Next = &ListNode{}
		currentNode = currentNode.Next
	}

	for list2 != nil {
		currentNode.Val = list2.Val
		list2 = list2.Next
		previousNode = currentNode
		currentNode.Next = &ListNode{}
		currentNode = currentNode.Next
	}

	if previousNode != nil && currentNode.Val == 0 {
		previousNode.Next = nil
	}

	return head
}
