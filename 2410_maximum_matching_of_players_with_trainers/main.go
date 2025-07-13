package main

import "sort"

func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.SliceStable(players, func(i, j int) bool {
		return players[i] > players[j]
	})
	sort.SliceStable(trainers, func(i, j int) bool {
		return trainers[i] > trainers[j]
	})

	count := 0
	for i, j := 0, 0; i < len(players) && j < len(trainers); i++ {
		if players[i] > trainers[j] {
			continue
		}

		count++
		j++
	}

	return count
}
