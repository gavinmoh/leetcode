package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		if prev != nil {
			divisor := gcd(current.Val, prev.Val)
			prev.Next = &ListNode{
				Val:  divisor,
				Next: current,
			}
		}
		prev = current
		current = current.Next
	}

	return head
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
