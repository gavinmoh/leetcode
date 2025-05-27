package main

func differenceOfSums(n int, m int) int {
	num1, num2 := 0, 0

	for num := 1; num <= n; num++ {
		if num%m != 0 {
			num1 += num
		} else {
			num2 += num
		}
	}

	return num1 - num2
}
