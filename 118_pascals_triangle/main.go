package main

func Generate(numRows int) [][]int {
	var result [][]int
	for i := 1; i <= numRows; i++ {
		var row []int
		if i == 1 {
			row = append(row, 1)
			result = append(result, row)
			continue
		}

		previousRow := result[i-2]
		for j := 1; j <= i; j++ {
			if j == 1 || j == i {
				row = append(row, 1)
			} else {
				left := previousRow[j-2]
				right := previousRow[j-1]
				row = append(row, left+right)
			}
		}
		result = append(result, row)
	}
	return result
}
