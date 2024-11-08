package main

func rotateTheBox(box [][]byte) [][]byte {
	rows, cols := len(box), len(box[0])

	result := make([][]byte, cols)
	for row := 0; row < cols; row++ {
		result[row] = make([]byte, rows)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			result[col][rows-row-1] = box[row][col]
		}
	}

	for col := 0; col < rows; col++ {
		high, low := cols-1, cols-1
		for high >= 0 {
			if result[high][col] == '#' {
				for low >= high {
					if result[low][col] == '.' {
						result[high][col] = '.'
						result[low][col] = '#'
						break
					}
					low--
				}
			}
			if result[high][col] == '*' {
				low = high
			}

			high--
		}
	}

	return result
}
