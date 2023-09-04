package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func HasCycle(head *ListNode) bool {
	cycle := false
	// Use a map to keep track of nodes we've seen
	nodes := make(map[*ListNode]bool)

	for head != nil {
		// If we've seen this node before, we have a cycle
		if _, ok := nodes[head]; ok {
			cycle = true
			break
		}
		// Otherwise, add it to the map and move on
		nodes[head] = true
		head = head.Next
	}

	return cycle
}
