package main

import "math/rand"

type RandomizedSet struct {
	Nums map[int]bool
}

func Constructor() RandomizedSet {
	return RandomizedSet{Nums: make(map[int]bool)}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, exists := this.Nums[val]
	this.Nums[val] = true
	return !exists
}

func (this *RandomizedSet) Remove(val int) bool {
	exists := this.Nums[val]
	if exists {
		delete(this.Nums, val)
	}
	return exists
}

func (this *RandomizedSet) GetRandom() int {
	random := rand.Intn(len(this.Nums))
	i := 0
	for val := range this.Nums {
		if i == random {
			return val
		}
		i++
	}
	return -1
}
