package main

import "fmt"

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

func main() {
	fmt.Println(convert("PAYPALISHIRING", 4))
}

// PAYPALISHIRING
// numRows = 3
// space = numRows - 2
// P   A   H   N
// A P L S I I G
// Y   I   R

// PAYPALISHIRING
// numRows = 4
// P     I    N
// A   L S  I G
// Y A   H R
// P     I

// PAYPALISHIRING
// numRows = 5
// P       H
// A     S I
// Y   I   R
// P L     I G
// A       N

// PAYPALISHIRING
// numRows = 6
// P         R
// A       I I
// Y     H   N
// P   S     G
// A I
// L
