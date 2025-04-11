package main

import "fmt"

func countSymmetricIntegers(low int, high int) int {
	count := 0

	for num := low; num <= high; num++ {
		numStr := fmt.Sprintf("%d", num)
		n := len(numStr)
		if n%2 != 0 {
			continue
		}

		mid := n / 2
		leftSum, rightSum := 0, 0
		for i, c := range numStr {
			d := int(c - '0')
			if i < mid {
				leftSum += d
			} else {
				rightSum += d
			}
		}

		if leftSum == rightSum {
			count++
		}
	}

	return count
}
