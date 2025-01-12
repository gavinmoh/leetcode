package main

func doesValidArrayExist(derived []int) bool {
	xorSum := 0
	for _, num := range derived {
		xorSum ^= num
	}

	return xorSum == 0
}
