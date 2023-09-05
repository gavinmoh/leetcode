package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func CopyRandomList(head *Node) *Node {
	var newHead *Node
	nodes := make(map[*Node]*Node)
	unknownRandom := false

	currentNode := head
	var currentNewNode *Node

	// go through the list and create new nodes
	for currentNode != nil {
		newNode := &Node{Val: currentNode.Val}
		// add the new current node to the map
		nodes[currentNode] = newNode

		if newHead == nil {
			newHead = newNode
			currentNewNode = newNode
		} else {
			currentNewNode.Next = newNode
			currentNewNode = currentNewNode.Next
		}

		if currentNode.Random != nil {
			// check if the random node is in the map
			if random, ok := nodes[currentNode.Random]; ok {
				// if it is, set the random node of the new node
				newNode.Random = random
			} else {
				// if not, set the unknownRandom flag to true
				// so that we can go through the list again
				unknownRandom = true
			}
		}

		// move to the next node
		currentNode = currentNode.Next
	}

	// if there are no unknown random nodes,
	// we know the list is complete and we can return
	if !unknownRandom {
		return newHead
	}

	// reset the current node to the head
	currentNode = head
	currentNewNode = newHead
	for currentNode != nil {
		if currentNode.Random != nil && currentNewNode.Random == nil {
			// set the random node of the new node
			currentNewNode.Random = nodes[currentNode.Random]
		}
		currentNode = currentNode.Next
		currentNewNode = currentNewNode.Next
	}

	return newHead
}
