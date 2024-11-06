package main

func minEnd(n int, x int) int64 {
	result := int64(x)

	for i := 0; i < n-1; i++ {
		result = (result + 1) | int64(x)
	}

	return result
}
