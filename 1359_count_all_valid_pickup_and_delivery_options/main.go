package main

const Modulo = 1000000007

func CountOrders(n int) int {
	if n == 1 {
		return 1
	}
	// this is a distributive property math problem
	// where the formula is n * (2n - 1) * (n - 1)!
	// where n is the number of orders
	// (2n - 1) is the number of possible positions for the first order
	// (n - 1)! is the number of possible positions for the remaining orders
	orders := n * (2*n - 1) * CountOrders(n-1)
	return orders % Modulo
}
