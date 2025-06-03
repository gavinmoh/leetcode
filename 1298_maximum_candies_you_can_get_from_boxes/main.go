package main

func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	n := len(status)
	queue := []int{}
	queue = append(queue, initialBoxes...)
	foundKeys := make([]bool, n)
	closedBoxes := map[int]bool{}

	result := 0
	// bfs
	for len(queue) > 0 {
		newQueue := []int{}
		for _, box := range queue {
			// check if box is open
			if status[box] == 0 {
				// check if we have the key
				if !foundKeys[box] {
					closedBoxes[box] = true
					continue
				}
			}

			// now the box is either open or we have the key
			// take all the candies
			result += candies[box]

			// take all the keys
			for _, key := range keys[box] {
				foundKeys[key] = true
			}

			// add all the boxes into the queue
			newQueue = append(newQueue, containedBoxes[box]...)
		}

		// check all the keys against closed boxes
		for closedBox := range closedBoxes {
			if foundKeys[closedBox] {
				newQueue = append(newQueue, closedBox)
				delete(closedBoxes, closedBox)
			}
		}
		queue = newQueue
	}

	return result
}
