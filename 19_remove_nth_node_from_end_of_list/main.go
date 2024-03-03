package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	current := head
	var prev *ListNode
	right := 1

	for current != nil {
		current = current.Next
		right++
		if right-n > 1 {
			if prev != nil {
				prev = prev.Next
			} else {
				prev = head
			}
		}
	}

	if prev == nil {
		return head.Next
	}

	if right-n == 1 {
		return nil
	}

	prev.Next = prev.Next.Next

	return head
}
