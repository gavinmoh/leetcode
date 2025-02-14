package main

type ProductOfNumbers struct {
	PrefixSum  []int
	RunningSum int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		PrefixSum:  []int{},
		RunningSum: 1,
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.RunningSum = 1
		this.PrefixSum = []int{}
	} else {
		this.RunningSum *= num
		this.PrefixSum = append(this.PrefixSum, this.RunningSum)
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	n := len(this.PrefixSum)
	if k > n {
		return 0
	}

	denominator := 1
	if k < n {
		denominator = this.PrefixSum[n-k-1]
	}

	return this.RunningSum / denominator
}

/**
* Your ProductOfNumbers object will be instantiated and called as such:
* obj := Constructor();
* obj.Add(num);
* param_2 := obj.GetProduct(k);
 */
