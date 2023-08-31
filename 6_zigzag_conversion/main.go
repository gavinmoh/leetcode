package main

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	var cols [][]string
	vertical := true
	verticalCount := numRows
	horizontalCount := numRows - 2

	for i := 0; i < len(s); {

		if vertical {
			var row []string
			for j := 0; j < verticalCount; j++ {
				if i < len(s) {
					row = append(row, string(s[i]))
				} else {
					row = append(row, "")
				}
				i++
			}
			cols = append(cols, row)

			// switch to horizontal
			vertical = false
		} else {
			for j := horizontalCount; j > 0; j-- {
				var row []string
				for k := 0; k < numRows; k++ {
					if k == j && i < len(s) {
						row = append(row, string(s[i]))
						i++
					} else {
						row = append(row, "")
					}
				}
				cols = append(cols, row)
			}

			// switch to vertical
			vertical = true
		}
	}

	// pick out all the characters horizontally
	result := ""

	for i := 0; i < numRows; i++ {
		for j := 0; j < len(cols); j++ {
			char := cols[j][i]
			if char != "" {
				result += char
			}
		}
	}

	return result
}
