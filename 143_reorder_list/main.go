package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	// use slow & fast to find the mid point of the list
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// reverse the second half of the list
	second := slow.Next
	slow.Next = nil
	var prev *ListNode
	for second != nil {
		temp := second.Next
		second.Next = prev
		prev = second
		second = temp
	}

	// merge 2 lists
	first, second := head, prev
	for second != nil {
		temp1, temp2 := first.Next, second.Next
		first.Next, second.Next = second, temp1
		first, second = temp1, temp2
	}
}
