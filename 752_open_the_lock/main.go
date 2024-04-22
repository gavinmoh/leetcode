package main

func openLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}

	deadendsHash := make(map[string]bool)
	for _, deadend := range deadends {
		deadendsHash[deadend] = true
	}

	if deadendsHash["0000"] || deadendsHash[target] {
		return -1
	}

	visited := make(map[string]bool)

	queue := []string{"0000"}
	for i := 1; len(queue) > 0; i++ {
		newQueue := []string{}
		for _, combo := range queue {
			if visited[combo] || deadendsHash[combo] {
				continue
			}

			// mark as visited
			visited[combo] = true

			// looping through all the position
			for _, pos := range []int{0, 1, 2, 3} {
				for _, val := range []int{1, -1} {
					newCombo := []rune(combo)
					num := newCombo[pos] + rune(val)
					if num > '9' {
						num = '0'
					} else if num < '0' {
						num = '9'
					}
					newCombo[pos] = num
					newComboStr := string(newCombo)

					if newComboStr == target {
						return i
					}

					newQueue = append(newQueue, newComboStr)
				}
			}
		}
		queue = newQueue
	}

	return -1
}
