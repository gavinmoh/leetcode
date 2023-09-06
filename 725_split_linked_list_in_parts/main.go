package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func SplitListToParts(head *ListNode, k int) []*ListNode {
	split := make([]*ListNode, k)
	nodes := make(map[int]*ListNode)
	nodesCount := 0

	for head != nil {
		nodesCount++
		nodes[nodesCount] = head
		head = head.Next
	}

	eachNode := nodesCount / k
	remainder := 0
	if eachNode < 1 {
		eachNode = 1
		remainder = 0
	} else {
		remainder = nodesCount % k
	}

	partsCount := 0
	for i := 1; i <= nodesCount; {
		var size int
		if remainder > 0 {
			size = eachNode + 1
			remainder--
		} else {
			size = eachNode
		}

		// adding the head node into the split
		split[partsCount] = nodes[i]
		partsCount++

		// breaking the list
		position := i + size - 1
		if position < nodesCount {
			nodes[position].Next = nil
		}

		nextPosition := i + size
		i = nextPosition
	}

	for partsCount < k {
		split[partsCount] = nil
		partsCount++
	}

	return split
}
