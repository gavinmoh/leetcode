package main

func hammingWeight(num uint32) int {
	var count int
	for num > 0 {
		if num&1 == 1 {
			count++
		}
		num >>= 1
	}
	return count
}
