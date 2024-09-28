package main

type MyCircularDeque struct {
	Data  []int
	Size  int
	Count int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		Data:  []int{},
		Size:  k,
		Count: 0,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.Count == this.Size {
		return false
	}

	this.Data = append([]int{value}, this.Data...)
	this.Count++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.Count == this.Size {
		return false
	}

	this.Data = append(this.Data, value)
	this.Count++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.Count == 0 {
		return false
	}

	this.Data = this.Data[1:]
	this.Count--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.Count == 0 {
		return false
	}

	this.Data = this.Data[:this.Count-1]
	this.Count--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.Count == 0 {
		return -1
	}

	return this.Data[0]
}

func (this *MyCircularDeque) GetRear() int {
	if this.Count == 0 {
		return -1
	}

	return this.Data[this.Count-1]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.Count == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.Count == this.Size
}

/**
* Your MyCircularDeque object will be instantiated and called as such:
* obj := Constructor(k);
* param_1 := obj.InsertFront(value);
* param_2 := obj.InsertLast(value);
* param_3 := obj.DeleteFront();
* param_4 := obj.DeleteLast();
* param_5 := obj.GetFront();
* param_6 := obj.GetRear();
* param_7 := obj.IsEmpty();
* param_8 := obj.IsFull();
 */
