package main

func passThePillow(n int, time int) int {
	i, direction := 1, 1

	for time > 0 {
		i += direction
		if i == n || i == 1 {
			direction *= -1
		}
		time--
	}

	return i
}
