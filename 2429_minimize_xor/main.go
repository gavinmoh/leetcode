package main

func minimizeXor(num1 int, num2 int) int {
	result := 0
	currentCount := 0
	target := setBitsCount(num2)
	currentBit := 31

	for currentCount < target {
		if isSet(num1, currentBit) || (target-currentCount > currentBit) {
			result = setBit(result, currentBit)
			currentCount++
		}
		currentBit--
	}

	return result
}

func setBitsCount(x int) int {
	count := 0
	for x > 0 {
		if x&1 == 1 {
			count++
		}
		x >>= 1
	}

	return count
}

func isSet(x, bit int) bool {
	bitMask := 1 << bit
	return x&bitMask != 0
}

func setBit(x, bit int) int {
	bitMask := 1 << bit
	return x | bitMask
}
