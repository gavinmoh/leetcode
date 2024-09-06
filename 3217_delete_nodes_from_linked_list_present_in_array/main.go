package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	numsMap := make(map[int]bool)
	for _, num := range nums {
		numsMap[num] = true
	}

	nodes := []*ListNode{}
	current := head
	for current != nil {
		if _, ok := numsMap[current.Val]; !ok {
			nodes = append(nodes, current)
		}

		current = current.Next
	}

	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	nodes[len(nodes)-1].Next = nil

	return nodes[0]
}
