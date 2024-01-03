package main

func numberOfBeams(bank []string) int {
	lasers := []int{}

	for _, str := range bank {
		laserCount := 0
		for _, char := range str {
			if char == '1' {
				laserCount++
			}
		}
		if laserCount > 0 {
			lasers = append(lasers, laserCount)
		}
	}

	laserBeams := 0
	for i := 1; i < len(lasers); i++ {
		laserBeams += lasers[i] * lasers[i-1]
	}

	return laserBeams
}
