package main

func minOperations(nums []int, k int) int {
	// calculate the xor of all the numbers in nums
	xor := 0
	for _, num := range nums {
		xor ^= num
	}

	// compare each bit of k and xor
	// if they're not equal, then we have to flip one bit
	count := 0
	for k > 0 || xor > 0 {
		if k%2 != xor%2 {
			count++
		}
		k /= 2
		xor /= 2
	}

	return count
}
