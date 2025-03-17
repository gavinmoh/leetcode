package main

func divideArray(nums []int) bool {
	freq := map[int]int{}
	for _, num := range nums {
		freq[num] ^= num
	}

	for _, xor := range freq {
		if xor != 0 {
			return false
		}
	}

	return true
}
