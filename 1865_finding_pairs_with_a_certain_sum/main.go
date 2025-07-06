package main

import "sort"

type FindSumPairs struct {
	nums1 []int
	nums2 []int
	freq  map[int]int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	sort.Ints(nums1)

	freq := map[int]int{}
	for _, num := range nums2 {
		freq[num]++
	}

	return FindSumPairs{
		nums1: nums1,
		nums2: nums2,
		freq:  freq,
	}
}

func (this *FindSumPairs) Add(index int, val int) {
	old := this.nums2[index]
	this.nums2[index] += val

	this.freq[old]--
	this.freq[this.nums2[index]]++
}

func (this *FindSumPairs) Count(tot int) int {
	result := 0
	for _, num := range this.nums1 {
		if num > tot {
			break
		}
		remain := tot - num
		if count, ok := this.freq[remain]; ok {
			result += count
		}
	}

	return result
}

/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */
