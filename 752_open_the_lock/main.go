package main

func openLock(deadends []string, target string) int {
	deadendsHash := make(map[string]bool)
	for _, deadend := range deadends {
		deadendsHash[deadend] = true
	}

	if deadendsHash["0000"] || deadendsHash[target] {
		return -1
	}

	visited := make(map[string]bool)

	queue := []string{"0000"}
	for i := 0; len(queue) > 0; i++ {
		newQueue := []string{}
		for _, combo := range queue {
			if combo == target {
				return i
			}

			if visited[combo] || deadendsHash[combo] {
				continue
			}

			// mark as visited
			visited[combo] = true

			// looping through all the position
			for j := 0; j < 4; j++ {
				// increment
				newQueue = append(newQueue, increment(combo, j))
				// descrement
				newQueue = append(newQueue, decrement(combo, j))
			}
		}
		queue = newQueue
	}

	return -1
}

func increment(combo string, pos int) string {
	newCombo := []rune(combo)
	num := newCombo[pos] + 1
	if num > '9' {
		num = '0'
	}
	newCombo[pos] = num

	return string(newCombo)
}

func decrement(combo string, pos int) string {
	newCombo := []rune(combo)
	num := newCombo[pos] - 1
	if num < '0' {
		num = '9'
	}
	newCombo[pos] = num

	return string(newCombo)
}
