package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var firstIndex int
	var secondIndex int

	for index, number := range nums {
		for i, n := range nums[index+1:] {
			if number+n == target {
				firstIndex = index
				secondIndex = i + index + 1
				break
			}
		}
	}
	return []int{firstIndex, secondIndex}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
