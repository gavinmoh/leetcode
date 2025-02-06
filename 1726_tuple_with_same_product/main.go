package main

func tupleSameProduct(nums []int) int {
	products := map[int]int{}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			product := nums[i] * nums[j]
			products[product]++
		}
	}

	result := 0
	for _, count := range products {
		pairs := (count - 1) * count / 2
		result += pairs * 8
	}

	return result
}
