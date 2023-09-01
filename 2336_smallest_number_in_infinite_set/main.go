package main

type Node struct {
	value int
	next  *Node
}

type SmallestInfiniteSet struct {
	head *Node
}

func Constructor() SmallestInfiniteSet {
	headNode := Node{value: 1, next: nil}
	currentNode := &headNode
	for i := 2; i <= 1000; i++ {
		newNode := Node{value: i, next: nil}
		currentNode.next = &newNode
		currentNode = &newNode
	}
	return SmallestInfiniteSet{head: &headNode}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	smallest := this.head.value
	this.head = this.head.next
	return smallest
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	newNode := Node{value: num, next: nil}

	currentNode := this.head
	previousNode := this.head

	for currentNode != nil {
		// skip if value already in set
		if currentNode.value == num {
			return
		}

		// if num is greater than current node,
		// check if next node is nil and add to end of list if so
		// otherwise, continue to next node
		if num > currentNode.value {
			if currentNode.next == nil {
				currentNode.next = &newNode
				return
			}
			previousNode = currentNode
			currentNode = currentNode.next
			continue
		}

		// if num is less than current node,
		// check if current node is head and add to beginning of list if so
		// otherwise, link previous node to new node and new node to current node
		if num < currentNode.value {
			if currentNode == this.head {
				newNode.next = this.head
				this.head = &newNode
				return
			} else {
				previousNode.next = &newNode
				newNode.next = currentNode
				return
			}
		}

		previousNode = currentNode
		currentNode = currentNode.next
	}
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
