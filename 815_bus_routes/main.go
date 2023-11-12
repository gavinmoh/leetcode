package main

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	// converting routes into hash table
	routeMaps := make(map[int]map[int]bool, len(routes))
	for i, route := range routes {
		routeMap := make(map[int]bool, len(route))
		for _, stop := range route {
			routeMap[stop] = true
		}
		routeMaps[i] = routeMap
	}

	queue := []int{}
	visited := make(map[int]bool, len(routes))
	for i, routeMap := range routeMaps {
		if routeMap[source] {
			queue = append(queue, i)
		}
	}

	busCount := 0
	// BFS
	for len(queue) > 0 {
		targetFound := false
		newQueue := []int{}
		// each time we enter a new loop, we are taking another bus
		busCount++
		for _, index := range queue {
			if visited[index] {
				continue
			}
			visited[index] = true

			routeMap := routeMaps[index]
			if routeMap[target] {
				targetFound = true
				break
			}

			for stop := range routeMap {
				// find the unvisited routes which contains this stop
				for i, routeMap := range routeMaps {
					if visited[i] {
						continue
					}
					if routeMap[stop] {
						newQueue = append(newQueue, i) // append it into new queue
					}
				}
			}
		}

		if targetFound {
			return busCount
		}
		// next loop we will look into next level of routes
		queue = newQueue
	}

	return -1
}
