package main

type MyQueue struct {
	Data []int
}

func Constructor() MyQueue {
	return MyQueue{Data: []int{}}
}

func (this *MyQueue) Push(x int) {
	this.Data = append(this.Data, x)
}

func (this *MyQueue) Pop() int {
	popVal := this.Data[0]
	this.Data = this.Data[1:]
	return popVal
}

func (this *MyQueue) Peek() int {
	return this.Data[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.Data) == 0
}

/**
* Your MyQueue object will be instantiated and called as such:
* obj := Constructor();
* obj.Push(x);
* param_2 := obj.Pop();
* param_3 := obj.Peek();
* param_4 := obj.Empty();
 */
