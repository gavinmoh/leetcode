package main

import "sort"

func findWinners(matches [][]int) [][]int {
	players := make(map[int][]int)

	for _, match := range matches {
		winner := match[0]
		loser := match[1]
		if _, ok := players[winner]; ok {
			players[winner][0]++
		} else {
			players[winner] = []int{1, 0}
		}

		if _, ok := players[loser]; ok {
			players[loser][1]++
		} else {
			players[loser] = []int{0, 1}
		}
	}

	allWin := []int{}
	lostOne := []int{}

	for player, result := range players {
		wins := result[0]
		loses := result[1]

		if wins+loses == 0 {
			continue
		}

		if loses == 0 {
			allWin = append(allWin, player)
		} else if loses == 1 {
			lostOne = append(lostOne, player)
		}
	}

	sort.Ints(allWin)
	sort.Ints(lostOne)

	return [][]int{allWin, lostOne}
}
