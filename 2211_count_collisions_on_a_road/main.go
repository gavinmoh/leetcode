package main

import "strings"

func countCollisions(directions string) int {
	directions = strings.TrimLeft(directions, "L")
	directions = strings.TrimRight(directions, "R")

	count := 0
	for _, c := range directions {
		if c == 'S' {
			continue
		}

		count++
	}

	return count
}
