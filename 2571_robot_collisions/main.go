package main

import "sort"

type Robot struct {
	Index     int
	Position  int
	Health    int
	Direction byte
}

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	robots := []*Robot{}
	for i := 0; i < len(positions); i++ {
		robot := &Robot{
			Index:     i,
			Position:  positions[i],
			Health:    healths[i],
			Direction: directions[i],
		}
		robots = append(robots, robot)
	}

	sort.SliceStable(robots, func(i, j int) bool {
		return robots[i].Position < robots[j].Position
	})

	stack := []*Robot{}
	for _, robot := range robots {
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			// collission only occurs in this scenario:
			// top -> <- robot
			if top.Direction == 'R' && robot.Direction == 'L' {
				// pop the top because only winner can be added back to the stack
				stack = stack[:len(stack)-1]

				if top.Health == robot.Health {
					robot = nil // remove both robots
					break
				}

				if top.Health > robot.Health {
					top.Health--
					robot = top
				} else {
					robot.Health--
				}
			} else {
				break
			}
		}
		if robot != nil {
			stack = append(stack, robot)
		}
	}

	sort.SliceStable(stack, func(i, j int) bool {
		return stack[i].Index < stack[j].Index
	})

	results := []int{}
	for _, robot := range stack {
		results = append(results, robot.Health)
	}

	return results
}
