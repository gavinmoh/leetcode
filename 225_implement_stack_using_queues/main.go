package main

type Node struct {
	Val  int
	Next *Node
}

type MyStack struct {
	Root *Node
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	newNode := &Node{Val: x, Next: this.Root}
	this.Root = newNode
}

func (this *MyStack) Pop() int {
	popNode := this.Root
	this.Root = this.Root.Next
	return popNode.Val
}

func (this *MyStack) Top() int {
	return this.Root.Val
}

func (this *MyStack) Empty() bool {
	return this.Root == nil
}

/**
* Your MyStack object will be instantiated and called as such:
* obj := Constructor();
* obj.Push(x);
* param_2 := obj.Pop();
* param_3 := obj.Top();
* param_4 := obj.Empty();
 */
