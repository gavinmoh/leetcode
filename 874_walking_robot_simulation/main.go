package main

type Direction int

const (
	NORTH Direction = iota
	SOUTH
	EAST
	WEST
)

func robotSim(commands []int, obstacles [][]int) int {
	direction := NORTH

	obstaclesMap := make(map[int]map[int]bool)
	for _, obstacle := range obstacles {
		x, y := obstacle[0], obstacle[1]
		if _, ok := obstaclesMap[x]; !ok {
			obstaclesMap[x] = make(map[int]bool)
		}
		obstaclesMap[x][y] = true
	}

	maxDistance := 0
	x, y := 0, 0
	for _, command := range commands {
		if command == -1 {
			direction = turnRight(direction)
		} else if command == -2 {
			direction = turnLeft(direction)
		} else {
			for i := 0; i < command; i++ {
				nextX, nextY := move(x, y, direction)
				if _, ok := obstaclesMap[nextX]; ok && obstaclesMap[nextX][nextY] {
					break
				}
				x, y = nextX, nextY
			}
			distance := x*x + y*y
			if distance > maxDistance {
				maxDistance = distance
			}
		}
	}

	return maxDistance
}

func turnLeft(direction Direction) Direction {
	switch direction {
	case NORTH:
		return WEST
	case SOUTH:
		return EAST
	case EAST:
		return NORTH
	case WEST:
		return SOUTH
	}

	return direction
}

func turnRight(direction Direction) Direction {
	switch direction {
	case NORTH:
		return EAST
	case SOUTH:
		return WEST
	case EAST:
		return SOUTH
	case WEST:
		return NORTH
	}

	return direction
}

func move(x int, y int, direction Direction) (int, int) {
	switch direction {
	case NORTH:
		y++
	case SOUTH:
		y--
	case EAST:
		x++
	case WEST:
		x--
	}

	return x, y
}
