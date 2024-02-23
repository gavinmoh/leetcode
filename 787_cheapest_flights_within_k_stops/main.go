package main

type Flight struct {
	dst   int
	price int
}

type City struct {
	flights []Flight
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	cities := make([]City, n)
	for i := 0; i < n; i++ {
		cities[i] = City{flights: []Flight{}}
	}

	for _, flight := range flights {
		source, destination, price := flight[0], flight[1], flight[2]
		newFlight := Flight{dst: destination, price: price}
		cities[source].flights = append(cities[source].flights, newFlight)
	}

	cache := make(map[int]map[int]map[int]int)
	var bfs func(int, int, int) int
	bfs = func(source, destination, stops int) int {
		if stops < 0 {
			return -1
		}

		if _, ok := cache[source]; !ok {
			cache[source] = make(map[int]map[int]int)
		}

		if _, ok := cache[source][destination]; !ok {
			cache[source][destination] = make(map[int]int)
		}

		if cachedResult, ok := cache[source][destination][stops]; ok {
			return cachedResult
		}

		costs := []int{}
		for _, flight := range cities[source].flights {
			if flight.dst == destination {
				costs = append(costs, flight.price)
			} else {
				cost := bfs(flight.dst, destination, stops-1)
				if cost != -1 {
					costs = append(costs, cost+flight.price)
				}
			}
		}

		result := -1
		if len(costs) > 0 {
			cheapest := costs[0]
			for i := 1; i < len(costs); i++ {
				cheapest = min(cheapest, costs[i])
			}

			result = cheapest
		}

		cache[source][destination][stops] = result
		return result
	}

	return bfs(src, dst, k)
}
