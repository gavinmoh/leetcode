package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func modifiedList(nums []int, head *ListNode) *ListNode {
	numsMap := make(map[int]bool)
	for _, num := range nums {
		numsMap[num] = true
	}

	// find new head
	for head != nil {
		if _, ok := numsMap[head.Val]; !ok {
			break
		}
		head = head.Next
	}

	tail, current := head, head.Next
	for current != nil {
		if _, ok := numsMap[current.Val]; !ok {
			tail.Next = current
			tail = current
		}

		current = current.Next
	}
	tail.Next = nil

	return head
}
