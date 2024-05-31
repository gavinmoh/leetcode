package main

func singleNumber(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}

	// find the position of rightmost 1 bit
	i := 0
	for xor&(1<<i) == 0 {
		i++
	}

	// find out the first number
	// we just need to xor the numbers
	// where ith bit is either 1 or 0
	firstNumber := 0
	for _, num := range nums {
		if num&(1<<i) == 0 {
			firstNumber ^= num
		}
	}

	// the second number can XOR out from xor using first number
	return []int{firstNumber, xor ^ firstNumber}
}
