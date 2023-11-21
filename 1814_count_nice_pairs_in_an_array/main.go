package main

func countNicePairs(nums []int) int {
	// tranform nums into nums[i] = nums[i] - rev(nums[i])
	for i, num := range nums {
		nums[i] = num - reverse(num)
	}

	valueMap := make(map[int]int, len(nums))
	output := 0

	modulo := 1_000_000_007

	for _, num := range nums {
		output = (output + valueMap[num]) % modulo
		valueMap[num] += 1
	}

	return output
}

// taken from 7 Reverse Integer
func reverse(x int) int {
	var result int
	for x != 0 {
		remainder := x % 10
		result = (result * 10) + remainder
		x /= 10
	}

	// Check for overflow
	if result > 2147483647 || result < -2147483648 {
		return 0
	}

	return result
}
