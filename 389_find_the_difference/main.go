package main

func findTheDifference(s string, t string) byte {
	// using XOR operator on each byte of the strings
	// all the same characters will cancel each other out
	// leaving the extra character
	var result byte
	for i := 0; i < len(s); i++ {
		result ^= s[i]
	}
	for i := 0; i < len(t); i++ {
		result ^= t[i]
	}

	return result
}
