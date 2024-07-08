package main

func findTheWinner(n int, k int) int {
	result := 0
	for i := 2; i < n+1; i++ {
		result = (result + k) % i
	}
	return result + 1
}
