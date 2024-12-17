package main

func finalPrices(prices []int) []int {
	result := make([]int, len(prices))
	for i, price := range prices {
		discount := 0
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= price {
				discount = prices[j]
				break
			}
		}
		result[i] = price - discount
	}

	return result
}
