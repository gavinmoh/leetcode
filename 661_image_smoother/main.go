package main

func imageSmoother(img [][]int) [][]int {
	rows := len(img)
	cols := len(img[0])
	output := make([][]int, rows)

	for row := 0; row < rows; row++ {
		output[row] = make([]int, cols)

		for col := 0; col < cols; col++ {
			cells := []int{img[row][col]}

			// top
			if row > 0 {
				cells = append(cells, img[row-1][col])

				// top left
				if col > 0 {
					cells = append(cells, img[row-1][col-1])
				}

				// top right
				if col < cols-1 {
					cells = append(cells, img[row-1][col+1])
				}
			}

			// bottom
			if row < rows-1 {
				cells = append(cells, img[row+1][col])

				// bottom left
				if col > 0 {
					cells = append(cells, img[row+1][col-1])
				}

				// bottom right
				if col < cols-1 {
					cells = append(cells, img[row+1][col+1])
				}
			}

			// left
			if col > 0 {
				cells = append(cells, img[row][col-1])
			}

			// right
			if col < cols-1 {
				cells = append(cells, img[row][col+1])
			}

			sum := 0
			for _, cell := range cells {
				sum += cell
			}

			output[row][col] = sum / len(cells)
		}
	}

	return output

}
