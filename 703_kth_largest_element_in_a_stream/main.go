package main

import "sort"

type KthLargest struct {
	K    int
	Data []int
}

func Constructor(k int, nums []int) KthLargest {
	return KthLargest{
		K:    k,
		Data: nums,
	}
}

func (this *KthLargest) Add(val int) int {
	this.Data = append(this.Data, val)
	sort.Ints(this.Data)
	n := len(this.Data)

	return this.Data[n-this.K]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
