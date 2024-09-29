package main

type Node struct {
	Freq int
	Prev *Node
	Next *Node
	Keys map[string]bool
}

type AllOne struct {
	Map  map[string]*Node
	Head *Node
	Tail *Node
}

func Constructor() AllOne {
	obj := AllOne{
		Map:  make(map[string]*Node),
		Head: &Node{Freq: 0},
		Tail: &Node{Freq: 0},
	}
	obj.Head.Next = obj.Tail
	obj.Tail.Prev = obj.Head

	return obj
}

func (this *AllOne) Inc(key string) {
	if node, ok := this.Map[key]; ok {
		freq := node.Freq
		delete(node.Keys, key)

		nextNode := node.Next
		if nextNode == this.Tail || nextNode.Freq != freq+1 {
			newNode := &Node{Freq: freq + 1, Keys: map[string]bool{key: true}}
			newNode.Prev = node
			newNode.Next = nextNode
			node.Next = newNode
			nextNode.Prev = newNode
			this.Map[key] = newNode
		} else {
			nextNode.Keys[key] = true
			this.Map[key] = nextNode
		}

		if len(node.Keys) == 0 {
			this.RemoveNode(node)
		}
	} else {
		firstNode := this.Head.Next
		if firstNode == this.Tail || firstNode.Freq > 1 {
			newNode := &Node{Freq: 1, Keys: map[string]bool{key: true}}
			newNode.Prev = this.Head
			newNode.Next = firstNode
			this.Head.Next = newNode
			firstNode.Prev = newNode
			this.Map[key] = newNode
		} else {
			firstNode.Keys[key] = true
			this.Map[key] = firstNode
		}
	}
}

func (this *AllOne) Dec(key string) {
	if _, ok := this.Map[key]; !ok {
		return
	}

	node := this.Map[key]
	delete(node.Keys, key)
	freq := node.Freq

	if freq == 1 {
		delete(this.Map, key)
	} else {
		prevNode := node.Prev
		if prevNode == this.Head || prevNode.Freq != freq-1 {
			newNode := &Node{Freq: freq - 1, Keys: map[string]bool{key: true}}
			newNode.Prev = prevNode
			newNode.Next = node
			node.Prev = newNode
			prevNode.Next = newNode
			this.Map[key] = newNode
		} else {
			prevNode.Keys[key] = true
			this.Map[key] = prevNode
		}
	}

	if len(node.Keys) == 0 {
		this.RemoveNode(node)
	}
}

func (this *AllOne) GetMaxKey() string {
	if this.Tail.Prev != this.Head {
		for key, _ := range this.Tail.Prev.Keys {
			return key
		}
	}
	return ""
}

func (this *AllOne) GetMinKey() string {
	if this.Head.Next != this.Tail {
		for key, _ := range this.Head.Next.Keys {
			return key
		}
	}

	return ""
}

func (this *AllOne) RemoveNode(node *Node) {
	prevNode := node.Prev
	nextNode := node.Next

	prevNode.Next = nextNode
	nextNode.Prev = prevNode
}

/**
* Your AllOne object will be instantiated and called as such:
* obj := Constructor();
* obj.Inc(key);
* obj.Dec(key);
* param_3 := obj.GetMaxKey();
* param_4 := obj.GetMinKey();
 */
