package main

func maximumInvitations(favorite []int) int {
	n := len(favorite)
	reversedGraph := make([][]int, n)
	for i := 0; i < n; i++ {
		reversedGraph[i] = []int{}
	}

	for person, favoritePerson := range favorite {
		reversedGraph[favoritePerson] = append(reversedGraph[favoritePerson], person)
	}

	bfs := func(start int, visited map[int]bool) int {
		queue := [][]int{{start, 0}}
		maxDistance := 0
		for len(queue) > 0 {
			// pop left
			currentNode, currentDistance := queue[0][0], queue[0][1]
			queue = queue[1:]

			for _, neighbour := range reversedGraph[currentNode] {
				if _, ok := visited[neighbour]; ok {
					continue
				}
				visited[neighbour] = true
				queue = append(queue, []int{neighbour, currentDistance + 1})
				maxDistance = max(maxDistance, currentDistance+1)
			}
		}
		return maxDistance
	}

	longestCycle, twoCycleInvitations := 0, 0
	visited := make([]bool, n)

	// find all cycles in graph
	for person := 0; person < n; person++ {
		if !visited[person] {
			visitedPeople := map[int]int{}
			currentPerson := person
			distance := 0

			for {
				if visited[currentPerson] {
					break
				}

				visited[currentPerson] = true
				visitedPeople[currentPerson] = distance
				distance++
				nextPerson := favorite[currentPerson]

				// cycle detected
				if _, ok := visitedPeople[nextPerson]; ok {
					cycleLength := distance - visitedPeople[nextPerson]
					longestCycle = max(longestCycle, cycleLength)

					// handle cycles with length of 2
					if cycleLength == 2 {
						visitedNodes := map[int]bool{
							currentPerson: true,
							nextPerson:    true,
						}
						twoCycleInvitations += (2 + bfs(nextPerson, visitedNodes) + bfs(currentPerson, visitedNodes))
					}

					break
				}
				currentPerson = nextPerson
			}
		}
	}

	return max(longestCycle, twoCycleInvitations)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
