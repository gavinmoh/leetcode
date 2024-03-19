package main

import "sort"

func leastInterval(tasks []byte, n int) int {
	frequencies := make([]int, 26)
	for _, task := range tasks {
		frequencies[task-'A']++
	}

	sort.Ints(frequencies)
	maxFreq := frequencies[25] - 1
	idleSlots := maxFreq * n

	for i := 24; i >= 0; i-- {
		idleSlots -= min(maxFreq, frequencies[i])
	}

	if idleSlots > 0 {
		return idleSlots + len(tasks)
	}

	return len(tasks)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
