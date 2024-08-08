package main

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	var isInMatrix = func(row, col int) bool {
		if row < 0 || col < 0 {
			return false
		}

		if row >= rows || col >= cols {
			return false
		}

		return true
	}

	nextDirection := map[rune]rune{
		'l': 'd',
		'd': 'r',
		'r': 'u',
		'u': 'l',
	}

	total := rows * cols
	result := [][]int{{rStart, cStart}}
	direction := 'l'
	steps := 1
	for len(result) < total {
		for i := 0; i < 2; i++ {
			for j := 0; j < steps; j++ {
				switch direction {
				case 'l':
					cStart += 1
				case 'd':
					rStart += 1
				case 'r':
					cStart -= 1
				case 'u':
					rStart -= 1
				}
				if isInMatrix(rStart, cStart) {
					result = append(result, []int{rStart, cStart})
				}
			}
			direction = nextDirection[direction]
		}
		steps += 1
	}

	return result
}
