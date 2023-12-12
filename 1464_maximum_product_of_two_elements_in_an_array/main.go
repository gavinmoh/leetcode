package main

func maxProduct(nums []int) int {
	largest := 0
	secondLargest := 0
	for _, num := range nums {
		if num > largest {
			secondLargest = largest
			largest = num
		} else {
			secondLargest = max(secondLargest, num)
		}
	}

	return (largest - 1) * (secondLargest - 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
