package main

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	edges := make([][]int, n)
	for i := 0; i < n-1; i++ {
		edges[i] = []int{i + 1}
	}

	result := []int{}

	for _, edge := range queries {
		from, to := edge[0], edge[1]
		edges[from] = append(edges[from], to)

		// bfs
		visited := map[int]bool{}
		cities := edges[0]
		distance := 1
	outer:
		for len(cities) > 0 {
			nextCities := []int{}
			for _, city := range cities {
				if city == n-1 {
					result = append(result, distance)
					break outer
				}

				if _, ok := visited[city]; ok {
					continue
				}
				visited[city] = true

				nextCities = append(nextCities, edges[city]...)
			}
			cities = nextCities
			distance++
		}
	}

	return result
}
