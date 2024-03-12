package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	newHead := &ListNode{Val: 0, Next: head}
	current := newHead

	for current != nil {
		prefixSum := 0
		next := current.Next

		for next != nil {
			prefixSum += next.Val
			if prefixSum == 0 {
				current.Next = next.Next
			}
			next = next.Next
		}
		current = current.Next
	}

	return newHead.Next
}
