package main

func numberOfPairs(points [][]int) int {
	count := 0
	for _, pointA := range points {
		x1, y1 := pointA[0], pointA[1]
		for _, pointB := range points {
			x2, y2 := pointB[0], pointB[1]
			if x1 == x2 && y1 == y2 {
				continue
			}

			// point A has to be upper left
			if x2 < x1 {
				continue
			}

			// point B has to lower than point A
			if y2 > y1 {
				continue
			}

			// check if any points between these A and B
			hasPointInBetween := false
			for _, pointC := range points {
				x3, y3 := pointC[0], pointC[1]
				if (x1 == x3 && y1 == y3) || (x2 == x3 && y2 == y3) {
					continue
				}

				if (x3 >= x1 && x3 <= x2) && (y3 >= y2 && y3 <= y1) {
					hasPointInBetween = true
					break
				}

			}
			if !hasPointInBetween {
				count++
			}
		}
	}

	return count
}
