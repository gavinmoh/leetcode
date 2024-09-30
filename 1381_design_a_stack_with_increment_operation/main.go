package main

type CustomStack struct {
	Data    []int
	Pointer int
	Size    int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		Data:    make([]int, maxSize),
		Pointer: -1,
		Size:    maxSize,
	}
}

func (this *CustomStack) Push(x int) {
	if this.Pointer < (this.Size - 1) {
		this.Pointer++
		this.Data[this.Pointer] = x
	}
}

func (this *CustomStack) Pop() int {
	if this.Pointer < 0 {
		return -1
	}

	val := this.Data[this.Pointer]
	this.Pointer--
	return val
}

func (this *CustomStack) Increment(k int, val int) {
	right := min(this.Pointer+1, k)
	for i := 0; i < right; i++ {
		this.Data[i] += val
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
* Your CustomStack object will be instantiated and called as such:
* obj := Constructor(maxSize);
* obj.Push(x);
* param_2 := obj.Pop();
* obj.Increment(k,val);
 */
